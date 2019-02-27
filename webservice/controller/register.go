package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"net/http"
)

//RegisterGET serves register page to users
func RegisterGET(w http.ResponseWriter, r *http.Request) {

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {

		// Check if the hash is a valid hash
		if course := model.CourseExists(hash); course.ID == -1 {
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
		v.Vars["Action"] = "/register"
	} else {
		v.Vars["Action"] = "/register?courseid=" + hash
	}
	v.Render(w)

	//todo check if there is a class hash in request
	//if there is, add the user logging in to the class and redirect
}

//RegisterPOST validates register requests from users
func RegisterPOST(w http.ResponseWriter, r *http.Request) {

	user := session.GetUserFromSession(r)

	if session.IsLoggedIn(r) { //already logged in, no need to register
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	name := r.FormValue("name")         // get form value name
	email := r.FormValue("email")       // get form value email
	password := r.FormValue("password") // get form value password
	hash := r.FormValue("courseid")     // get from link courseID

	//check that nothing is empty and password match passwordConfirm
	if name == "" || email == "" || password == "" || password != r.FormValue("passwordConfirm") { //login credentials cannot be empty
		http.Redirect(w, r, "/", http.StatusBadRequest) //400 bad request
		return
	}

	user, ok := model.RegisterUser(name, email, password) //register user in database

	if ok {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)
		// Add new user to course

		if hash != "" {
			if id := model.CourseExists(hash).ID; id != -1 {
				model.AddUserToCourse(user.ID, id)
			}
		}
	} else {
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		return
	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}
