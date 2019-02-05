package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){

	//check if user is logged in

	//check if there is a class id in request
	//if there is, add the user logging in to the class and redirect

	w.WriteHeader(http.StatusOK)

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
