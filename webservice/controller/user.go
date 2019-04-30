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
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"
)

// UserGET serves user page to users
func UserGET(w http.ResponseWriter, r *http.Request) {
	// Services
	services := service.NewServices(db.GetDB())

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Get courses to user
	courses, err := services.Course.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "user/profile"

	v.Vars["User"] = currentUser
	v.Vars["Courses"] = courses

	v.Render(w)
}

// UserUpdatePOST changes the user information
func UserUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()
	// Get current user
	currentUser := session.GetUserFromSession(r)
	// Services
	services := service.NewServices(db.GetDB())

	// Get hashed password
	hash, err := services.User.FetchHash(currentUser.ID)
	if err != nil {
		log.Println("services, user, fetch hash", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get new data from the form
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != currentUser.EmailPrivate.String {
		updatedUser := currentUser
		updatedUser.EmailPrivate = sql.NullString{
			String: p.Sanitize(secondaryEmail),
			Valid:  secondaryEmail != "",
		}

		// Check if the email exists in the db
		exist, _, err := services.User.EmailExists(p.Sanitize(secondaryEmail))
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// If the email already exists
		if exist {
			log.Println("email already exist in db")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		mailService := mail.Mail{}

		// Get new hash in 20 chars
		validationHash := xid.NewWithTime(time.Now()).String()

		// Fill validationEmail model for new insert in table
		validationEmail := model.ValidationEmail{
			Hash:      validationHash,
			TimeStamp: util.GetTimeInCorrectTimeZone(),
		}

		validationEmail.UserID = sql.NullInt64{
			Int64: int64(currentUser.ID),
			Valid: currentUser.ID != 0,
		}

		// Insert into the db
		validationID, err := services.Validation.Insert(validationEmail)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("EmailExists, ", err.Error())
			return
		}

		userData := model.UserRegistrationPending{
			Email:        secondaryEmail,
			ValidationID: validationID,
		}

		_, err = services.UserPending.InsertNewEmail(userData)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("insert users_pending, ", err.Error())
			return
		}

		// Get link
		baseURL := "http://" + r.Host
		link := baseURL + "/confirmemail?id=" + validationHash

		// Set subject and message
		subject := "Confirm new Secondary Email | CSAMS"
		message := "Hi " + currentUser.Name + ",\n\n" +
			"This email is sent by the CS Assignment Submission System.\n" +
			"We have received an request to add a new secondary email to your user profile on CSAMS (" + baseURL + ")\n" +
			"If you have not requested this and suspect a hacking attempt, please contact your lecturer.\n\n" +
			"Click this link to confirm your email:\n" +
			link

		// Send email with link
		err = mailService.SendSingleRecipient(userData.Email, subject, message)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("mail, confirm new email", err.Error())
			return
		}
	}

	// if no of these fields are empty, try to change password
	if oldPass != "" && newPass != "" && repeatPass != "" {

		// Check if the password is correct length
		if len(newPass) < 6 {
			ErrorHandler(w, r, http.StatusBadRequest)
			log.Println("Password is to short, minimum 6 chars in length!")
			return
		}

		// check if the old password is correct
		err = util.CompareHashAndPassword(oldPass, hash)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// If password hasn't changed or new and repeat doesn't match
		if newPass == oldPass || newPass != repeatPass {
			log.Println("New password has not changed or repeat password and new password is not the same")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Update password
		err = services.User.UpdatePassword(currentUser.ID, newPass)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Log password change in the database and give error if something went wrong
		err = services.Logs.InsertChangePassword(currentUser.ID)
		if err != nil {
			log.Println("log, change password ", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

	}

	//UserGET(w, r)
	http.Redirect(w, r, "/user", http.StatusFound) //success redirect to homepage
}

// ConfirmEmailGET validates new email requests from users
func ConfirmEmailGET(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("id")

	// Check that has is not empty
	if hash == "" {
		log.Println("hash can not be empty")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get currentUser
	currentUser := session.GetUserFromSession(r)

	// Services
	services := service.NewServices(db.GetDB())

	match, payload, err := services.Validation.Match(hash)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("hashMatch, ", err.Error())
		return
	}

	// Only do stuff if there is a match
	if match {
		timeNow := util.GetTimeInCorrectTimeZone()

		// Check if the link has expired (after 12 hours)
		if timeNow.After(payload.TimeStamp.Add(time.Hour * 12)) {
			// Update forgottenPass table to be expired (one time use only!)
			err = services.Validation.UpdateValidation(payload.ID, false)
			if err != nil {
				ErrorHandler(w, r, http.StatusInternalServerError)
				log.Println("validation, update validation, ", err.Error())
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

		// Get all pending users
		users, err := services.UserPending.FetchAll()
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("users_pending fetch all, ", err.Error())
			return
		}

		// Check for the user with the corresponding validation id
		newUser := model.UserRegistrationPending{}
		for _, user := range users {
			if user.ValidationID == payload.ID {
				newUser = *user
			}
		}

		// Update validation table to be expired (one time use only!)
		err = services.Validation.UpdateValidation(newUser.ValidationID, false)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("forgotten, update validation, ", err.Error())
			return
		}

		// Fetch user
		user, err := services.User.Fetch(int(payload.UserID.Int64))
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("services, user, fetch, ", err.Error())
			return
		}

		// Add new email
		user.EmailPrivate = sql.NullString{
			String: newUser.Email,
			Valid:  newUser.Email != "",
		}

		// Update user with new email
		err = services.User.Update(user.ID, *user)
		if err != nil {
			log.Println("services, user, update", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Log changes to db
		err = services.Logs.InsertChangeEmail(user.ID, currentUser.EmailPrivate.String, user.EmailPrivate.String)
		if err != nil {
			log.Println("log, change secondary email ", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		//update session
		currentUser.EmailPrivate = user.EmailPrivate
		session.SaveUserToSession(currentUser, w, r)

		session.SaveMessageToSession("Your secondary email is now updated", w, r)

	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage

}
