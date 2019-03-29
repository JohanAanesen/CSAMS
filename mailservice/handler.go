package main

import (
	"net/http"
)

// HandlerGET handles GET requests
func HandlerGET(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// HandlerPOST handles POST requests
func HandlerPOST(w http.ResponseWriter, r *http.Request) {


}
