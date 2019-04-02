package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

//Load loads the router handlers
func Load() http.Handler {
	return routes()
}

//LoadHTTP loads the router handler through http
func LoadHTTP() http.Handler {
	return routes()
}

//LoadHTTPS loads the router handler through https
func LoadHTTPS() http.Handler {
	return routes()
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://%s", r.Host), http.StatusMovedPermanently)
}

func routes() http.Handler {
	// Instantiate mux-router
	router := mux.NewRouter().StrictSlash(true)

	// Handlers for sending forgotten password email to single user
	router.HandleFunc("/", ForgottenPassGET).Methods("GET")
	router.HandleFunc("/", ForgottenPassPOST).Methods("POST")

	// Handlers for sending a single email to an user
	router.HandleFunc("/single", SingleMailGET).Methods("GET")
	router.HandleFunc("/single", SingleMailPOST).Methods("POST")

	// Handlers for sending one email to multiple users
	router.HandleFunc("/multiple", MultipleMailGET).Methods("GET")
	router.HandleFunc("/multiple", MultipleMailPOST).Methods("POST")

	return handlers.CombinedLoggingHandler(os.Stdout, router)
}
