package handlers

import "net/http"

func ClassHandler(w http.ResponseWriter, r *http.Request){

	//check if request has an classID

	//check if user is an participant of said class or a teacher

	//get info from db

	//parse info to html template

}

func ClassListHandler(w http.ResponseWriter, r *http.Request){

	//check if request has an classID

	//check if user is an participant of said class or a teacher

	//get classlist from db

	//parse info to html template

}