package handlers

import (
	"fmt"
	"net/http"
)

func AssignmentHandler(w http.ResponseWriter, r *http.Request){

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		fmt.Fprintf(w, "Id is %s\n", id)
		if id == "" {
			//redirect to error page
		}
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

func AssignmentAutoHandler(w http.ResponseWriter, r *http.Request){

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		fmt.Fprintf(w, "Id is %s\n", id)
		if id == "" {
			//redirect to error page
		}
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

func AssignmentPeerHandler(w http.ResponseWriter, r *http.Request){

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		fmt.Fprintf(w, "Id is %s\n", id)
		if id == "" {
			//redirect to error page
		}
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}
