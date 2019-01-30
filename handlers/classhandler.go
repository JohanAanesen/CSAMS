package handlers

import (
	"fmt"
	"net/http"
)

func ClassHandler(w http.ResponseWriter, r *http.Request){

	//check if request has an classID
	if r.Method == http.MethodGet{

		id := r.FormValue("id")

		if id == ""{
			//redirect to error page
			ErrorHandler(w, r, http.StatusForbidden)
			return
		}

		fmt.Fprintf(w, "Id is %s\n", id)

		//check if id is valid through database

		//check if user is an participant of said class or a teacher

	}


	//get info from db

	//parse info to html template

}

func ClassListHandler(w http.ResponseWriter, r *http.Request){

	//check if request has an classID
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		if id == "" {
			//redirect to error page
			ErrorHandler(w, r, http.StatusForbidden)
			return
		}

		fmt.Fprintf(w, "Id is %s\n", id)
	}
	//check if user is an participant of said class or a teacher

	//get classlist from db

	//parse info to html template

}