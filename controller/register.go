package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"net/http"
)

//RegisterGET serves register page to users
func RegisterGET(w http.ResponseWriter, r *http.Request) {
	if session.IsLoggedIn(r) {
		IndexGET(w, r)
		return
	}

	v := view.New(r)
	v.Name = "register"
	v.Render(w)

	//todo check if there is a class id in request
	//if there is, add the user logging in to the class and redirect
}

//RegisterPOST validates register requests from users
func RegisterPOST(w http.ResponseWriter, r *http.Request) {

	user := session.GetUserFromSession(r)

	if session.IsLoggedIn(r) { //already logged in, no need to register
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	name := r.FormValue("name")         //get form value name
	email := r.FormValue("email")       //get form value email
	password := r.FormValue("password") //get form value password

	//check that nothing is empty and password match passwordConfirm
	if name == "" || email == "" || password == "" || password != r.FormValue("passwordConfirm") { //login credentials cannot be empty
		http.Redirect(w, r, "/", http.StatusBadRequest) //400 bad request
		return
	}

	user, ok := db.RegisterUser(name, email, password) //register user in database

	if ok {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)
	} else {
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		return
	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}