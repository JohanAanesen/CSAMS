package controller

import (
	"encoding/gob"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/mail"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"
)

func init() {
	//todo maybe move this?
	gob.Register(model.User{})
}

//LoginGET serves login page to users
func LoginGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())

	// Models
	course := model.Course{}

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {
		// Check if course exists based on hash, and return it
		course = *courseService.Exists(hash)
		// Check if the hash is a valid hash
		if course.ID == -1 {
			log.Println("course id is -1")
			ErrorHandler(w, r, http.StatusBadRequest)
			hash = ""
			return
		}
	}

	// Check if user is already logged in
	currentUser := session.GetUserFromSession(r)
	if currentUser.Authenticated { //already logged in, redirect to homepage
		// If hash was valid, add user isn't in the course, then add user to course
		userInCourse, err := courseService.UserInCourse(currentUser.ID, course.ID)
		if err != nil {
			log.Println("course service, user in course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Check if user is not in course nad that hash is not blank
		if !userInCourse && hash != "" {
			// Add user to course
			err := courseService.AddUser(currentUser.ID, course.ID)
			if err == service.ErrUserAlreadyInCourse {
				http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
				return
			}

			// Check for errors
			if err != nil {
				log.Println("add user to course", err)
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}

		// Redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound) //redirect
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "login"

	// Send the correct link to template
	if hash == "" {
		v.Vars["Action"] = ""
	} else {
		v.Vars["Action"] = "?courseid=" + hash
	}

	v.Vars["Message"] = session.GetAndDeleteMessageFromSession(w, r)

	v.Render(w)
}

//LoginPOST validates login requests
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()
	// Get current user from session
	currentUser := session.GetUserFromSession(r)

	if currentUser.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	// Services
	userService := service.NewUserService(db.GetDB())
	courseService := service.NewCourseService(db.GetDB())

	email := r.FormValue("email")       // email
	password := r.FormValue("password") // password
	hash := r.FormValue("courseid")     // courseID from link

	if email == "" || password == "" { //login credentials cannot be empty
		session.SaveMessageToSession("Credentials cannot be empty!", w, r)
		LoginGET(w, r)
		return
	}

	user, err := userService.Authenticate(p.Sanitize(email), p.Sanitize(password))
	if err != nil {
		log.Println("user service authenticate", err)
		session.SaveMessageToSession("Wrong credentials!", w, r)
		LoginGET(w, r) //try again
		return
	}

	//save user to session values
	user.Authenticated = true
	session.SaveUserToSession(user, w, r)

	// Add new user to course, if he's not in the course
	if hash != "" {
		course := courseService.Exists(hash)
		userInCourse, err := courseService.UserInCourse(currentUser.ID, course.ID)
		if err != nil {
			log.Println("course service, user in course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		if course.ID != -1 && !userInCourse {
			err := courseService.AddUser(user.ID, course.ID)

			if err == service.ErrUserAlreadyInCourse {
				http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
				return
			}

			if err != nil {
				log.Println("add user to course", err)
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}
	}

	// Removes any messages saved earlier, in failed attempts
	session.GetAndDeleteMessageFromSession(w, r)
	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage //todo change redirection
	//IndexGET(w, r) //redirect to homepage
}

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
	userService := service.NewUserService(db.GetDB())
	forgottenService := service.NewForgottenPassService(db.GetDB())

	exists, userID, err := userService.EmailExists(email)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("EmailExists, ", err.Error())
		return
	}

	// If the email exists in the db
	if exists {
		// Get new hash in 20 chars
		hash := xid.NewWithTime(time.Now()).String()

		// Fill forgotten model for new insert in table
		forgotten := model.ForgottenPass{
			UserID:    userID,
			Hash:      hash,
			TimeStamp: util.GetTimeInCorrectTimeZone(),
		}

		// Insert into the db
		_, err := forgottenService.Insert(forgotten)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("EmailExists, ", err.Error())
			return
		}

		// Send email with link TODO send link, not hash
		mailservice := mail.Mail{}
		err = mailservice.SendMail(email, "http://localhost:8088/forgotpassword?id="+hash)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("mail.SendMail, ", err.Error())
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "forgotten"
	v.Vars["Message"] = "If the email provided exists, we will send you and email with instructions"
	v.Vars["Email"] = email

	v.Render(w)

}

// changePasswordPOST checks the hash and time, and changes password if it's correct
func changePasswordPOST(password string, hash string, w http.ResponseWriter, r *http.Request) {

	// Services
	//userService := service.NewUserService(db.GetDB())
	forgottenService := service.NewForgottenPassService(db.GetDB())

	match, payload, err := forgottenService.Match(hash)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("hashMatch, ", err.Error())
		return
	}

	if match {

		// Check if the link has expired (after 24 hours)
		if payload.TimeStamp.Add(time.Hour * 24).Before(util.GetTimeInCorrectTimeZone()) {
			ErrorHandler(w, r, http.StatusBadRequest)
			log.Println("Link expired")
			return
		}

		// TODO brede
		// Change password to user

		// Update forgottenPass table to be expired (one time use only!)

		// Give feedback
		log.Println("link not expired")
	} else {
		ErrorHandler(w, r, http.StatusBadRequest)
		log.Println("Link has no match in db")
		return
	}
}
