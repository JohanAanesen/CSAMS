package handlers

import (
	"../../db"
	"html/template"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	session, err := db.CookieStore.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	//check if user is already logged in

	if session.Values["username"] != ""{
		http.Redirect(w, r, "/", http.StatusFound)
	}

	//check if there is a class id in request
	//if there is, add the user logging in to the class and redirect

	//parse template

	if session.IsNew {
	}
		temp, err := template.ParseFiles("web/login.html")
		if err != nil {
			log.Fatal(err)
		}

		temp.Execute(w, nil)


}

func LoginRequest(w http.ResponseWriter, r *http.Request){
	session, err := db.CookieStore.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if r.FormValue("email") == "" || r.FormValue("password") == ""{
		return
	}

	if db.UserAuth(r.FormValue("email"), r.FormValue("password")) == false{
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	session.Values["username"] = r.FormValue("email")


	err = session.Save(r, w)
	if err != nil{
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}
