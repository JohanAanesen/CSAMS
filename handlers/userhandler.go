package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request){


	//check that user is logged in
		// code




	//fetch users information from server

	//parse information with template
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/user.html")
	if err != nil{
		log.Fatal(err)
	}

	temp.Execute(w, nil)

}
