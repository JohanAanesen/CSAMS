package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int){

	w.WriteHeader(status)

	if status == http.StatusForbidden{
		fmt.Fprint(w, "403 Forbidden")
	}else if status == http.StatusBadRequest{
		fmt.Fprint(w, "400 Bad Request")
	}else if status == http.StatusInternalServerError{
		fmt.Fprint(w, "500 Internal Server Error")
	}else if status == http.StatusNotFound{
		fmt.Fprint(w, "404 Not Found")
	}else if status == http.StatusNotImplemented{
		fmt.Fprint(w, "501 Not Implemented")
	}else if status == http.StatusUnauthorized{
		fmt.Fprint(w, "401 Unauthorized")
	}
	//todo: html and stuff, link back to homepage/where you were
}
