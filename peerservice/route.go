package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Load() http.Handler {
	return routes()
}

func LoadHTTP() http.Handler {
	return routes()
}

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
