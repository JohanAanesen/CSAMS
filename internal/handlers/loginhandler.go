package handlers

import (
	"../../db"
	"encoding/gob"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	ID 				int
	Email      		string
	Authenticated 	bool
}

func init(){
	gob.Register(User{})
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
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
	}{
		PageTitle: "Sign In",
	}); err != nil {
		log.Fatal(err)
	}


}

func LoginRequest(w http.ResponseWriter, r *http.Request){
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Fatal(err)//todo log this event
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	user := getUser(session)
	if user.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password") //password


	if email == "" || password == ""{  //login credentials cannot be empty
		return
	}

	userid, ok := db.UserAuth(email, password) //authenticate user

	if ok {
		session.Values["user"] = User{ID: userid, Email: email, Authenticated: true}
	}else{
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		fmt.Println(err.Error())
		return
	}

	err = session.Save(r, w)	//save data to session
	if err != nil{
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		fmt.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func RegisterRequest(w http.ResponseWriter, r *http.Request){
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}

	user := getUser(session)
	if user.Authenticated { //already logged in, no need to register
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == ""{  //login credentials cannot be empty
		return
	}

	userid, ok := db.RegisterUser(email, password) //register user in database

	if ok {
		session.Values["user"] = User{ID: userid, Email: email, Authenticated: true}
	}else{
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		return
	}

	err = session.Save(r, w)	//save data to session
	if err != nil{
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func getUser(s *sessions.Session) User{
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

