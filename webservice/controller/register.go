package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
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
		RegisterGET(w, r)
		return
	}

	// Services
	userService := service.NewUserService(db.GetDB())
	courseService := service.NewCourseService(db.GetDB())

	//Sanitize input
	name = p.Sanitize(name)
	email = p.Sanitize(email)
	password = p.Sanitize(password)

	userData := model.User{
		Name: name,
		EmailStudent: email,
	}

	userID, err := userService.Register(userData, password)
	//user, err := model.RegisterUser(name, email, password) //register user in database
	if err != nil {
		log.Println(err.Error())
		session.SaveMessageToSession("Email already in use!", w, r)
		RegisterGET(w, r)
		return
	}

	user, err := userService.Fetch(userID)
	if err != nil {
		log.Println("user service fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if user.ID != 0 {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(*user, w, r)
		// Add new user to course

		if hash != "" {
			hash = p.Sanitize(hash)
			course := courseService.Exists(hash)
			if course.ID != -1 {
				err = courseService.AddUser(user.ID, course.ID)
				if err != nil {
					log.Println("course service add user", err.Error())
					ErrorHandler(w, r, http.StatusInternalServerError)
					return
				}
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}
