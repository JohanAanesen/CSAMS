package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request){

	//check that error is provided in uri
	if r.Method == http.MethodGet{

		msg := r.FormValue("error")

		if msg != ""{
			fmt.Fprintf(w, "%s\n", msg)
		}else{
			fmt.Fprintln(w, "This is an error page, but we don't know what the error is sorry.")
		}

	}

	//todo: html and stuff, link back to homepage/where you were
}
