package route

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route/middleware/logrequest"
	"github.com/gorilla/mux"
	"net/http"
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

func routes() *mux.Router {
	// Instantiate mux-router
	router := mux.NewRouter().StrictSlash(true)

	// Set path prefix for the static-folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Middleware for logging Requests
	router.Use(logrequest.Handler)

	// Index-page Handlers
	router.HandleFunc("/", controller.IndexGET).Methods("GET")

	// Course-page Handlers
	// Example:
	// router.HandleFunc("/course", controller.CourseGET).Methods("GET")

	// Assignment-page Handlers

	// User-page Handlers

	// Dashboard-page Handlers

	return router
}