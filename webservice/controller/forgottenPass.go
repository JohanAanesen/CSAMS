package controller

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/mail"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"
)

// ForgottenGET serves the forgotten password page to students
func ForgottenGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "forgotten"

	// Clear message and email
	v.Vars["Message"] = ""
	v.Vars["Email"] = ""
	v.Vars["Hash"] = r.FormValue("id") // hash

	v.Render(w)
}

// ForgottenPOST checks routes the two different post requests
func ForgottenPOST(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")                 // email
	newPass := r.FormValue("newPassword")         // new password
	confirmPass := r.FormValue("confirmPassword") // confirm password
	hash := r.FormValue("id")

	// Route where the POST request are going
	if email != "" {
		sendEmailPOST(email, w, r)
	} else if newPass != "" && confirmPass != "" && newPass == confirmPass {

		// Check if the password is correct length
		if len(newPass) < 6 {
			ErrorHandler(w, r, http.StatusBadRequest)
			log.Println("Password is to short, minimum 6 chars in length!")
			return
		}

		changePasswordPOST(newPass, hash, w, r)
	} else {
		ErrorHandler(w, r, http.StatusBadRequest)
		log.Println("Something wrong with the credentials!")
		return
	}
}

// sendEmailPOST checks if the email is valid and sends a link to the email to change password
func sendEmailPOST(email string, w http.ResponseWriter, r *http.Request) {

	// Services
	services := service.NewServices(db.GetDB())

	exists, userID, err := services.User.EmailExists(email)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("EmailExists, ", err.Error())
		return
	}

	// If the email exists in the db
	if exists {
		// Get new hash in 20 chars
		hash := xid.NewWithTime(time.Now()).String()

		// Fill validationEmail model for new insert in table
		validationEmail := model.ValidationEmail{
			Hash:      hash,
			TimeStamp: util.GetTimeInCorrectTimeZone(),
		}

		validationEmail.UserID = sql.NullInt64{
			Int64: int64(userID),
			Valid: userID != 0,
		}

		// Insert into the db
		_, err := services.Validation.Insert(validationEmail)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("EmailExists, ", err.Error())
			return
		}

		// Get link
		link := "http://" + r.Host + "/forgotpassword?id=" + hash

		// Send email with link
		mailservice := mail.Mail{}
		err = mailservice.MailForgottenPassword(email, link)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("mail.MailForgottenPassword, ", err.Error())
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	v := view.New(r)
	v.Name = "forgotten"
	v.Vars["Message"] = "If the email provided exists, we will send you and email with instructions"
	v.Vars["Email"] = email

	v.Render(w)

}

// changePasswordPOST checks the hash and time, and changes password if it's correct
func changePasswordPOST(password string, hash string, w http.ResponseWriter, r *http.Request) {

	// Services
	services := service.NewServices(db.GetDB())

	match, payload, err := services.Validation.Match(hash)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("hashMatch, ", err.Error())
		return
	}

	// If the hash matches
	if match {

		timeNow := util.GetTimeInCorrectTimeZone()

		// Check if the link has expired (after 12 hours)
		if timeNow.After(payload.TimeStamp.Add(time.Hour * 12)) {
			// Update forgottenPass table to be expired (one time use only!)
			err = services.Validation.UpdateValidation(payload.ID, false)
			if err != nil {
				ErrorHandler(w, r, http.StatusInternalServerError)
				log.Println("forgotten, update validation, ", err.Error())
				return
			}

			ErrorHandler(w, r, http.StatusBadRequest)
			log.Println("Link expired, it's been 12 hours since creation")
			return
		}

		// Check if the link has been used before (one time use only)
		if !payload.Valid {
			ErrorHandler(w, r, http.StatusBadRequest)
			log.Println("Link expired, used before")
			return
		}

		// Change password to user
		err := services.User.UpdatePassword(int(payload.UserID.Int64), password)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("user, change password, ", err.Error())
			return
		}

		// Update forgottenPass table to be expired (one time use only!)
		err = services.Validation.UpdateValidation(payload.ID, false)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("forgotten, update validation, ", err.Error())
			return
		}

		// Log password change to db
		err = services.Logs.InsertChangePasswordEmail(int(payload.UserID.Int64))
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("log, user change password with email, ", err.Error())
			return
		}

		// Give feedback
		_ = session.SetFlash("You can now log in with your new password", w, r)
		http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage //todo change redirection

	} else {
		ErrorHandler(w, r, http.StatusBadRequest)
		log.Println("Link has no match in db")
		return
	}
}
