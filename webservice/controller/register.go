package controller

import (
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

//RegisterGET serves register page to users
func RegisterGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())

	name := r.FormValue("name")   // get form value name
	email := r.FormValue("email") // get form value email

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {
		course := courseService.Exists(hash)
		// Check if the hash is a valid hash
		if course.ID == -1 {
			log.Println("course id is -1")
			ErrorHandler(w, r, http.StatusBadRequest)
			hash = ""
			return
		}
	}

	if session.IsLoggedIn(r) {
		IndexGET(w, r)
		return
	}

	v := view.New(r)
	v.Name = "register"
	// Send the correct link to template
	if hash == "" {
		v.Vars["Action"] = ""
	} else {
		v.Vars["Action"] = "?courseid=" + hash
	}

	v.Vars["Name"] = name
	v.Vars["Email"] = email

	v.Vars["Message"] = session.GetAndDeleteMessageFromSession(w, r)

	v.Render(w)

	//todo check if there is a class hash in request
	//if there is, add the user logging in to the class and redirect
}

//RegisterPOST validates register requests from users
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	//XSS sanitizer
	p := bluemonday.UGCPolicy()

	currentUser := session.GetUserFromSession(r)

	if currentUser.Authenticated { //already logged in, no need to register
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	name := r.FormValue("name")         // get form value name
	email := r.FormValue("email")       // get form value email
	password := r.FormValue("password") // get form value password
	hash := r.FormValue("courseid")     // get from link courseID

	//check that nothing is empty and password match passwordConfirm
	if name == "" || email == "" || password == "" || password != r.FormValue("passwordConfirm") { //login credentials cannot be empty
		session.SaveMessageToSession("Passwords does not match or fields are empty!", w, r)
		log.Println("passwords does not match or fields are empty!")
		RegisterGET(w, r)
		return
	}

	// Check if the password is correct length
	if len(password) < 6 {
		ErrorHandler(w, r, http.StatusBadRequest)
		log.Println("Password is to short, minimum 6 chars in length!")
		return
	}

	// Services
	userService := service.NewUserService(db.GetDB())
	//courseService := service.NewCourseService(db.GetDB())
	userPendingService := service.NewUserPendingService(db.GetDB())
	validationService := service.NewValidationService(db.GetDB())

	mailService := mail.Mail{}

	//Sanitize input
	name = p.Sanitize(name)
	email = p.Sanitize(email)
	password = p.Sanitize(password)
	hash = p.Sanitize(hash)

	/*
		userData := model.User{
			Name:         name,
			EmailStudent: email,
		}
	*/

	userData := model.UserRegistrationPending{
		Name:         name,
		EmailStudent: email,
		Password:     password,
	}

	// get status if the email exists in db or not
	exist, _, err := userService.EmailExists(userData.EmailStudent)
	if err != nil {
		log.Println(err.Error())
		RegisterGET(w, r)
		return
	}

	// Check if email already exist in db
	if exist {
		log.Println("Email already exists")
		session.SaveMessageToSession("Email already in use", w, r)
		RegisterGET(w, r)
		return
	}

	// Get new hash in 20 chars
	validationHash := xid.NewWithTime(time.Now()).String()

	// Fill forgotten model for new insert in table
	forgotten := model.ValidationEmail{
		Hash:      validationHash,
		TimeStamp: util.GetTimeInCorrectTimeZone(),
	}

	// Insert into the db
	validationID, err := validationService.Insert(forgotten)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("EmailExists, ", err.Error())
		return
	}

	userData.ValidationID = validationID

	_, err = userPendingService.Insert(userData)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("insert users_pending, ", err.Error())
		return
	}

	// Get link TODO brede : add real link here with possible courseid hash
	baseURL := "http://" + r.Host
	link := baseURL + "/confirm?id=" + validationHash

	// Set subject and message
	subject := "Confirm new User"
	message := "Hi " + userData.Name + ",\n\n" +
		"There has been requested to create an user on CSAMS (" + baseURL + ")\n" +
		"If this was not you, please disregard this email.\n\n" +
		"Click this link to confirm your email:\n" +
		link

	// Send email with link
	err = mailService.SendSingleRecipient(userData.EmailStudent, subject, message)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println("mail.MailForgottenPassword, ", err.Error())
		return
	}

	session.SaveMessageToSession("Please confirm your email first", w, r)

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}

// ConfirmGET validates register requests from users
func ConfirmGET(w http.ResponseWriter, r *http.Request) {

	hash := r.FormValue("id")

	// Check that has is not empty
	if hash == "" {
		log.Println("hash can not be empty")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Services
	forgottenService := service.NewValidationService(db.GetDB())
	userPendingService := service.NewUserPendingService(db.GetDB())
	userService := service.NewUserService(db.GetDB())
	//logService := service.NewLogsService(db.GetDB())

	match, payload, err := forgottenService.Match(hash)
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
			err = forgottenService.UpdateValidation(payload.ID, false)
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

		// Get all pending users
		users, err := userPendingService.FetchAll()
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

		// Update forgottenPass table to be expired (one time use only!)
		err = forgottenService.UpdateValidation(newUser.ValidationID, false)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("forgotten, update validation, ", err.Error())
			return
		}

		// Create new user
		user := model.User{
			Name:         newUser.Name,
			EmailStudent: newUser.EmailStudent,
			Teacher:      false,
		}

		// Get correct password from user
		password, err := userPendingService.FetchPassword(newUser.ID)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("userpending, fetch password, ", err.Error())
			return
		}

		// RegisterWithHashing user (insert to database)
		_, err = userService.Register(user, password)
		if err != nil {
			log.Println(err.Error())
			RegisterGET(w, r)
			return
		}

		/* TODO brede : uncomment this and test
		// Log new user to db
		err = logService.InsertNewUser(newUser.ID)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("new user log", err.Error())
			return
		}
		*/

		session.SaveMessageToSession("You can now log in with your email and password", w, r)

	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}
