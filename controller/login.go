package controller

import (
	"encoding/gob"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
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

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {

		course = db.CourseExists(hash)

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
		if hash != "" && !db.UserExistsInCourse(user.ID, course.ID) {
			db.AddUserToCourse(user.ID, course.ID)
			// TODO : maybe redirect to course page ?
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "login"

	// Send the correct link to template
	if hash == "" {
		v.Vars["Action"] = "/login"
	} else {
		v.Vars["Action"] = "/login?courseid=" + hash
	}

	v.Render(w)
}

//LoginPOST validates login requests
func LoginPOST(w http.ResponseWriter, r *http.Request) {

	user := session.GetUserFromSession(r)

	if user.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	email := r.FormValue("email")       // email
	password := r.FormValue("password") // password
	hash := r.FormValue("courseid")     // courseID from link

	if email == "" || password == "" { //login credentials cannot be empty
		LoginGET(w, r)
		return
	}

	user, ok := db.UserAuth(email, password) //authenticate user

	if ok {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)

		// Add new user to course, if he's not in the course
		if hash != "" {
			if id := db.CourseExists(hash).ID; id != -1 && !db.UserExistsInCourse(user.ID, id) {
				db.AddUserToCourse(user.ID, id)
				// TODO : maybe redirect to course page ?
			}
		}

	} else {
		//redirect to errorhandler
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		log.Println("LoginPOST error")
		return
	}

	//http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage //todo change redirection
	IndexGET(w, r) //redirect to homepage
}
