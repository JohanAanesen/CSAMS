package handlers

import (
	"encoding/gob"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

func init() {
	gob.Register(model.User{})
}

//LoginHandler serves login page to users
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError) //error getting session 500
		return
	}

	//check if user is already logged in
	user := getUser(session)
	if user.Authenticated { //already logged in, redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//todo check if there is a class id in request
	//if there is, add the user logging in to the class and redirect

	//parse template
	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/login.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
	}{
		PageTitle: "Sign In",
		Menu:      util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}

}

//LoginRequest validates login requests
func LoginRequest(w http.ResponseWriter, r *http.Request) {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	user := getUser(session)
	if user.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password") //password

	if email == "" || password == "" { //login credentials cannot be empty
		LoginHandler(w,r)
		return
	}

	user, ok := db.UserAuth(email, password) //authenticate user

	if ok {
		//save user to session values
		user.Authenticated = true
		session.Values["user"] = user
	} else {
		//redirect to errorhandler
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	err = session.Save(r, w) //save session changes

	if err != nil {
		log.Println(err.Error())//todo log this event
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage
	MainHandler(w, r) //redirect to homepage
}

func getUser(s *sessions.Session) model.User {
	val := s.Values["user"]
	var user = model.User{}
	user, ok := val.(model.User)
	if !ok {
		return model.User{Authenticated: false}
	}
	return user
}
