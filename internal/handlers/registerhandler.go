package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"html/template"
	"log"
	"net/http"
)

//RegisterHandler serves register page to users
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	//check if user is logged in
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
	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/register.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle   string
		LoadFormCSS bool
		Menu        page.Menu
	}{
		PageTitle:   "Sign Up",
		LoadFormCSS: true,
		Menu:        util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
}

//RegisterRequest validates register requests from users
func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		log.Fatal(err)
		return
	}

	user := getUser(session)
	if user.Authenticated { //already logged in, no need to register
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

	userid, name, ok := db.RegisterUser(name, email, password) //register user in database

	if ok {
		//save user to session values
		session.Values["user"] = User{ID: userid, Name: name, Email: email, Authenticated: true}
	} else {
		ErrorHandler(w, r, http.StatusUnauthorized)
		//todo log this event
		return
	}

	err = session.Save(r, w) //save session changes
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Fatal(err)
		//todo log this event
		return
	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}
