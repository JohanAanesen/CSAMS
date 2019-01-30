package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request){

	//send user to login if no valid login cookies exist

	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/test.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
