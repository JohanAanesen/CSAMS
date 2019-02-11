package handlers

import (
	"fmt"
	"net/http"
)

//UserHandler serves user page to users
func UserHandler(w http.ResponseWriter, r *http.Request) {

	//check that user is logged in

	//fetch users information from server

	//parse information with template
	fmt.Fprintf(w, "this makes the test go through")
}
