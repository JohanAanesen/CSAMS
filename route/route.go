package route

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// Load ... TODO (Svein) add comment here
func Load() http.Handler {
	return routes()
}

// LoadHTTPS ... TODO (Svein): Add TLS settings
func LoadHTTPS() http.Handler {
	return routes()
}

// redirectToHTTPS ....
func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://%s", r.Host), http.StatusMovedPermanently)
}

// routes setups all routes
func routes() http.Handler {
	// Instantiate mux-router
	router := mux.NewRouter().StrictSlash(true)

	// Index-page Handlers
	router.HandleFunc("/", controller.IndexGET).Methods("GET")
	router.HandleFunc("/", controller.JoinCoursePOST).Methods("POST")

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
	adminrouter := router.PathPrefix("/admin").Subrouter()
	adminrouter.Use(middleware.TeacherAuth)

	adminrouter.HandleFunc("/", controller.AdminGET).Methods("GET")

	adminrouter.HandleFunc("/course", controller.AdminCourseGET).Methods("GET")
	adminrouter.HandleFunc("/course/create", controller.AdminCreateCourseGET).Methods("GET")
	adminrouter.HandleFunc("/course/create", controller.AdminCreateCoursePOST).Methods("POST")
	adminrouter.HandleFunc("/course/update/{id}", controller.AdminUpdateCourseGET).Methods("GET")
	adminrouter.HandleFunc("/course/update/{id}", controller.AdminUpdateCoursePOST).Methods("POST")

	adminrouter.HandleFunc("/assignment", controller.AdminAssignmentGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/create", controller.AdminAssignmentCreateGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/create", controller.AdminAssignmentCreatePOST).Methods("POST")

	adminrouter.HandleFunc("/submission", controller.AdminSubmissionGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreateGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreatePOST).Methods("POST")

	adminrouter.HandleFunc("/faq", controller.AdminFaqGET).Methods("GET")
	adminrouter.HandleFunc("/faq/edit", controller.AdminFaqEditGET).Methods("GET")
	adminrouter.HandleFunc("/faq/update", controller.AdminFaqUpdatePOST).Methods("POST")

	adminrouter.HandleFunc("/settings", controller.AdminSettingsGET).Methods("GET")
	adminrouter.HandleFunc("/settings", controller.AdminSettingsPOST).Methods("POST")
	adminrouter.HandleFunc("/settings/import", controller.AdminSettingsImportGET).Methods("GET")
	adminrouter.HandleFunc("/settings/import", controller.AdminSettingsImportPOST).Methods("POST")

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


	return handlers.CombinedLoggingHandler(os.Stdout, router)
}
