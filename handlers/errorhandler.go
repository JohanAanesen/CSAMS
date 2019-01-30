package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int){

	w.WriteHeader(status)

	if status == http.StatusForbidden{
		fmt.Fprint(w, "403 Forbidden")
	}
	//todo: html and stuff, link back to homepage/where you were
}
