package route

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/go-chi/chi/middleware"
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

	// go-chi/chi/middleware/logger.go
	router.Use(middleware.Logger)

	// Index-page Handlers
	router.HandleFunc("/", controller.IndexGET).Methods("GET")

	// Course-page Handlers
	router.HandleFunc("/course", controller.CourseGET).Methods("GET")
	router.HandleFunc("/course/list", controller.CourseListGET).Methods("GET")

	// Assignment-page Handlers
	router.HandleFunc("/assignment", controller.AssignmentGET).Methods("GET")
	router.HandleFunc("/assignment/peer", controller.AssignmentPeerGET).Methods("GET")
	router.HandleFunc("/assignment/auto", controller.AssignmentAutoGET).Methods("GET")

	// User-page Handlers
	router.HandleFunc("/user", controller.UserGET).Methods("GET")
	router.HandleFunc("/user/update", controller.UserUpdatePOST).Methods("POST")

	// Admin-page Handlers
	router.HandleFunc("/admin", controller.AdminGET).Methods("GET")
	router.HandleFunc("/admin/course", controller.AdminCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/create", controller.AdminCreateCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/create", controller.AdminCreateCoursePOST).Methods("POST")
	router.HandleFunc("/admin/course/update/{id}", controller.AdminUpdateCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/update/{id}", controller.AdminUpdateCoursePOST).Methods("POST")
	router.HandleFunc("/admin/assignment", controller.AdminAssignmentGET).Methods("GET")

	// Login/Register Handlers
	router.HandleFunc("/login", controller.LoginGET).Methods("GET")
	router.HandleFunc("/login", controller.LoginPOST).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")
	router.HandleFunc("/logout", controller.LogoutGET).Methods("GET")

	// Set path prefix for the static-folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// 404 Error Handler
	router.NotFoundHandler = http.HandlerFunc(controller.NotFoundHandler)
	// 405 Error Handler
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedHandler)

	return router
}
