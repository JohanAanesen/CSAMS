package controller

import (
	"encoding/gob"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"html/template"
	"log"
	"net/http"
)

func init() {
	//todo maybe move this?
	gob.Register(model.User{})
}

//LoginGET serves login page to users
func LoginGET(w http.ResponseWriter, r *http.Request) {
	//check if user is already logged in
	user := session.GetUserFromSession(r)
	if user.Authenticated { //already logged in, redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//todo check if there is a class id in request
	//if there is, add the user logging in to the class and redirect

	//parse template
	temp, err := template.ParseFiles("template/layout.html", "template/navbar.html", "template/login.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      model.Menu
	}{
		PageTitle: "Sign In",
		Menu:      util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}

}

//LoginPOST validates login requests
func LoginPOST(w http.ResponseWriter, r *http.Request) {

	user := session.GetUserFromSession(r)
	if user.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password") //password

	if email == "" || password == "" { //login credentials cannot be empty
		LoginGET(w, r)
		return
	}

	user, ok := db.UserAuth(email, password) //authenticate user

	if ok {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)
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