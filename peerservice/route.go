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

	// Index-page Handlers
	router.HandleFunc("/", HandlerGET).Methods("GET")
	router.HandleFunc("/", HandlerPOST).Methods("POST")

	return handlers.CombinedLoggingHandler(os.Stdout, router)
}
