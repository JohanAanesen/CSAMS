package handlers

import (
	"html/template"
	"log"
	"net/http"
	"../../db"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request){

	//check if user is logged in
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

	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/register.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
	}{
		PageTitle: "Sign Up",
	}); err != nil {
		log.Fatal(err)
	}
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

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if name == "" || email == "" || password == "" || password != r.FormValue("passwordConfirm"){  //login credentials cannot be empty
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	userid, name, ok := db.RegisterUser(name, email, password) //register user in database

	if ok {
		session.Values["user"] = User{ID: userid, Name: name, Email: email, Authenticated: true}
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