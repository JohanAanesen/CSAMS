package controller

import (
	"encoding/gob"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
)

func init() {
	//todo maybe move this?
	gob.Register(model.User{})
}

//LoginGET serves login page to users
func LoginGET(w http.ResponseWriter, r *http.Request) {

	course := model.Course{}

	//course repo
	courseRepo := &model.CourseRepository{}

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {

		course = courseRepo.CourseExists(hash)

		// Check if the hash is a valid hash
		if course.ID == -1 {
			ErrorHandler(w, r, http.StatusBadRequest)
			hash = ""
			return
		}
	}

	// Check if user is already logged in
	user := session.GetUserFromSession(r)
	if user.Authenticated { //already logged in, redirect to homepage

		// If hash was valid, add user isn't in the course, then add user to course
		if hash != "" && !courseRepo.UserExistsInCourse(user.ID, course.ID) {
			courseRepo.AddUserToCourse(user.ID, course.ID)
		}

		http.Redirect(w, r, "/", http.StatusFound) //redirect
		return
	}

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.WriteHeader(http.StatusOK)

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

	//sanitizer
	p := bluemonday.UGCPolicy()

	user := session.GetUserFromSession(r)

	if user.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	email := r.FormValue("email")       // email
	password := r.FormValue("password") // password
	hash := r.FormValue("courseid")     // courseID from link

	if email == "" || password == "" { //login credentials cannot be empty
		session.SaveMessageToSession("Credentials cannot be empty!", w, r)
		LoginGET(w, r)
		return
	}

	user, ok := model.UserAuth(p.Sanitize(email), p.Sanitize(password)) //authenticate user

	//course repo
	courseRepo := &model.CourseRepository{}

	if ok {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)

		// Add new user to course, if he's not in the course
		if hash != "" {
			if id := courseRepo.CourseExists(hash).ID; id != -1 && !courseRepo.UserExistsInCourse(user.ID, id) {
				courseRepo.AddUserToCourse(user.ID, id)
			}
		}

	} else {
		//redirect to errorhandler //todo return message to user and let them login again
		session.SaveMessageToSession("Wrong credentials!", w, r)
		LoginGET(w, r) //try again
		//todo log this event
		log.Println("LoginPOST error")
		return
	}

	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage //todo change redirection
	//IndexGET(w, r) //redirect to homepage
}

func ForgottenGET(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "forgotten"

	v.Render(w)
}
