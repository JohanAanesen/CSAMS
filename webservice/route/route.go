package route

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/route/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// Load http handler
func Load() http.Handler {
	return routes()
}

// LoadHTTPS http handler
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
	router.HandleFunc("/assignment/{id:[0-9]+}", controller.AssignmentSingleGET).Methods("GET")
	router.HandleFunc("/assignment/peer", controller.AssignmentPeerGET).Methods("GET")
	router.HandleFunc("/assignment/auto", controller.AssignmentAutoGET).Methods("GET")
	router.HandleFunc("/assignment/submission", controller.AssignmentUploadGET).Methods("GET")
	router.HandleFunc("/assignment/submission/update", controller.AssignmentUploadPOST).Methods("POST")
	router.HandleFunc("/assignment/{id:[0-9]+}/submission/{userid:[0-9]+}", controller.AssignmentUserSubmissionGET).Methods("GET")
	router.HandleFunc("/assignment/{id:[0-9]+}/submission/{userid:[0-9]+}", controller.AssignmentUserSubmissionPOST).Methods("POST")

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
	adminrouter.HandleFunc("/course/update/{id:[0-9]+}", controller.AdminUpdateCourseGET).Methods("GET")
	adminrouter.HandleFunc("/course/update", controller.AdminUpdateCoursePOST).Methods("POST")

	adminrouter.HandleFunc("/course/{id:[0-9]+}/assignments", controller.AdminCourseAllAssignments).Methods("GET")

	adminrouter.HandleFunc("/assignment", controller.AdminAssignmentGET).Methods("GET")

	adminrouter.HandleFunc("/assignment/{id:[0-9]+}", controller.AdminSingleAssignmentGET).Methods("GET")

	adminrouter.HandleFunc("/assignment/create", controller.AdminAssignmentCreateGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/create", controller.AdminAssignmentCreatePOST).Methods("POST")

	adminrouter.HandleFunc("/assignment/update/{id:[0-9]+}", controller.AdminUpdateAssignmentGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/update", controller.AdminUpdateAssignmentPOST).Methods("POST")

	adminrouter.HandleFunc("/assignment/{id:[0-9]+}/submissions", controller.AdminAssignmentSubmissionsGET).Methods("GET")
	//adminrouter.HandleFunc("/assignment/{id:[0-9]+}/submission", controller.AdminAssignmentSubmissionGET).Methods("GET")

	adminrouter.HandleFunc("/submission", controller.AdminSubmissionGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreateGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreatePOST).Methods("POST")
	adminrouter.HandleFunc("/submission/update/{id:[0-9]+}", controller.AdminSubmissionUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/submission/update", controller.AdminSubmissionUpdatePOST).Methods("POST")
	adminrouter.HandleFunc("/submission/delete", controller.AdminSubmissionDELETE).Methods("DELETE")

	adminrouter.HandleFunc("/review", controller.AdminReviewGET).Methods("GET")
	adminrouter.HandleFunc("/review/create", controller.AdminReviewCreateGET).Methods("GET")
	adminrouter.HandleFunc("/review/create", controller.AdminReviewCreatePOST).Methods("POST")
	adminrouter.HandleFunc("/review/update/{id:[0-9]+}", controller.AdminReviewUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/review/update", controller.AdminReviewUpdatePOST).Methods("POST")
	adminrouter.HandleFunc("/review/delete", controller.AdminReviewDELETE).Methods("DELETE")

	adminrouter.HandleFunc("/scheduler", controller.AdminSchedulerGET).Methods("GET")
	adminrouter.HandleFunc("/scheduler/delete", controller.AdminSchedulerDELETE).Methods("POST")

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
