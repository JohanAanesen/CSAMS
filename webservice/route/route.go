package route

import (
	"github.com/JohanAanesen/CSAMS/webservice/controller"
	"github.com/JohanAanesen/CSAMS/webservice/route/middleware"
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

// routes setups all routes
func routes() http.Handler {
	// Instantiate mux-router
	router := mux.NewRouter().StrictSlash(true)

	// User-page router
	userrouter := router.PathPrefix("/").Subrouter()
	userrouter.Use(middleware.UserAuth)

	// Admin-page router
	adminrouter := router.PathPrefix("/admin").Subrouter()
	adminrouter.Use(middleware.TeacherAuth)

	// Index-page Handlers
	userrouter.HandleFunc("/", controller.IndexGET).Methods("GET")
	userrouter.HandleFunc("/", controller.JoinCoursePOST).Methods("POST")

	// Course-page Handlers
	userrouter.HandleFunc("/course/{id:[0-9]+}", controller.CourseGET).Methods("GET")
	userrouter.HandleFunc("/course/list", controller.CourseListGET).Methods("GET") // TODO (Svein): Needed?

	// Assignment-page Handlers
	userrouter.HandleFunc("/assignment", controller.AssignmentGET).Methods("GET") // TODO (Svein): Needed?
	userrouter.HandleFunc("/assignment/{id:[0-9]+}", controller.AssignmentSingleGET).Methods("GET")
	userrouter.HandleFunc("/assignment/peer", controller.AssignmentPeerGET).Methods("GET") // TODO (Svein): Needed?
	userrouter.HandleFunc("/assignment/auto", controller.AssignmentAutoGET).Methods("GET") // TODO (Svein): Needed?
	userrouter.HandleFunc("/assignment/submission", controller.AssignmentUploadGET).Methods("GET")
	userrouter.HandleFunc("/assignment/submission/{id:[0-9]+}/withdraw", controller.AssignmentWithdrawGET).Methods("GET")
	userrouter.HandleFunc("/assignment/submission/update", controller.AssignmentUploadPOST).Methods("POST")
	userrouter.HandleFunc("/assignment/{id:[0-9]+}/submission/{userid:[0-9]+}", controller.AssignmentUserSubmissionGET).Methods("GET")
	userrouter.HandleFunc("/assignment/{id:[0-9]+}/submission/{userid:[0-9]+}", controller.AssignmentUserSubmissionPOST).Methods("POST")
	userrouter.HandleFunc("/assignment/{id:[0-9]+}/requestreview", controller.AssignmentReviewRequestPOST).Methods("POST")

	// User-page Handlers
	userrouter.HandleFunc("/user", controller.UserGET).Methods("GET")
	userrouter.HandleFunc("/user/update", controller.UserUpdatePOST).Methods("POST")

	// Admin-page Handlers
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
	//adminrouter.HandleFunc("/assignment/{id:[0-9]+}/submission", controller.AdminAssignmentSubmissionGET).Methods("GET") // TODO (Svein): Needed?
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}", controller.AdminAssignmentSingleSubmissionGET).Methods("GET")

	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}", controller.AdminAssignmentSingleSubmissionGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}/create", controller.AdminAssignmentSubmissionCreateGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}/create", controller.AdminAssignmentSubmissionCreatePOST).Methods("POST")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}/update", controller.AdminAssignmentSubmissionUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}/update", controller.AdminAssignmentSubmissionUpdatePOST).Methods("POST")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/submission/{userID:[0-9]+}", controller.AdminAssignmentSubmissionDELETE).Methods("DELETE")

	//adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/review/{userID:[0-9]+}", controller.AdminAssignmentReviewsGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/review/{targetID:[0-9]+}/{reviewerID:[0-9]+}/update", controller.AdminAssignmentReviewsUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/review/{targetID:[0-9]+}/{reviewerID:[0-9]+}/update", controller.AdminAssignmentReviewsUpdatePOST).Methods("POST")

	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/review", controller.AdminAssignmentReviewGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/review/{userID:[0-9]+}", controller.AdminAssignmentSingleReviewGET).Methods("GET")
	adminrouter.HandleFunc("/assignment/{assignmentID:[0-9]+}/reviews_done/{userID:[0-9]+}", controller.AdminAssignmentSingleReviewsDoneGET).Methods("GET")

	adminrouter.HandleFunc("/submission", controller.AdminSubmissionGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreateGET).Methods("GET")
	adminrouter.HandleFunc("/submission/create", controller.AdminSubmissionCreatePOST).Methods("POST")
	adminrouter.HandleFunc("/submission/update/{id:[0-9]+}", controller.AdminSubmissionUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/submission/update/{id:[0-9]+}/used", controller.AdminSubmissionUpdateUsedPOST).Methods("POST")
	adminrouter.HandleFunc("/submission/update/{id:[0-9]+}/weights", controller.AdminSubmissionUpdateWeightsGET).Methods("GET")
	adminrouter.HandleFunc("/submission/update/{id:[0-9]+}/weights", controller.AdminSubmissionUpdateWeightsPOST).Methods("POST")
	adminrouter.HandleFunc("/submission/update", controller.AdminSubmissionUpdatePOST).Methods("POST")
	adminrouter.HandleFunc("/submission/delete", controller.AdminSubmissionDELETE).Methods("DELETE")

	adminrouter.HandleFunc("/review", controller.AdminReviewGET).Methods("GET")
	adminrouter.HandleFunc("/review/create", controller.AdminReviewCreateGET).Methods("GET")
	adminrouter.HandleFunc("/review/create", controller.AdminReviewCreatePOST).Methods("POST")
	adminrouter.HandleFunc("/review/update/{id:[0-9]+}", controller.AdminReviewUpdateGET).Methods("GET")
	adminrouter.HandleFunc("/review/update/{id:[0-9]+}/used", controller.AdminReviewUpdateUsedPOST).Methods("POST")
	adminrouter.HandleFunc("/review/update/{id:[0-9]+}/weights", controller.AdminReviewUpdateWeightsGET).Methods("GET")
	adminrouter.HandleFunc("/review/update/{id:[0-9]+}/weights", controller.AdminReviewUpdateWeightsPOST).Methods("POST")
	adminrouter.HandleFunc("/review/update", controller.AdminReviewUpdatePOST).Methods("POST")
	adminrouter.HandleFunc("/review/delete", controller.AdminReviewDELETE).Methods("DELETE")

	adminrouter.HandleFunc("/changepass", controller.AdminChangePassGET).Methods("GET")
	adminrouter.HandleFunc("/changepass/list", controller.AdminGetUsersPOST).Methods("POST")

	adminrouter.HandleFunc("/faq", controller.AdminFaqGET).Methods("GET")
	adminrouter.HandleFunc("/faq/edit", controller.AdminFaqEditGET).Methods("GET")
	adminrouter.HandleFunc("/faq/update", controller.AdminFaqUpdatePOST).Methods("POST")

	adminrouter.HandleFunc("/settings", controller.AdminSettingsGET).Methods("GET")
	adminrouter.HandleFunc("/settings", controller.AdminSettingsPOST).Methods("POST")
	adminrouter.HandleFunc("/settings/import", controller.AdminSettingsImportGET).Methods("GET")
	adminrouter.HandleFunc("/settings/import", controller.AdminSettingsImportPOST).Methods("POST")

	// Login/RegisterWithHashing Handlers
	router.HandleFunc("/login", controller.LoginGET).Methods("GET")
	router.HandleFunc("/login", controller.LoginPOST).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")
	router.HandleFunc("/privacy", controller.PrivacyGET).Methods("GET")
	userrouter.HandleFunc("/logout", controller.LogoutGET).Methods("GET")

	// Login forgotten password handler
	router.HandleFunc("/forgotpassword", controller.ForgottenGET).Methods("GET")
	router.HandleFunc("/forgotpassword", controller.ForgottenPOST).Methods("POST")

	// RegisterWithHashing confirm user
	router.HandleFunc("/confirm", controller.ConfirmGET).Methods("GET")

	// Set path prefix for the static-folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// 404 Error Handler
	router.NotFoundHandler = middleware.UserAuth(http.HandlerFunc(controller.NotFoundHandler))
	// 405 Error Handler
	router.MethodNotAllowedHandler = middleware.UserAuth(http.HandlerFunc(controller.MethodNotAllowedHandler))

	return handlers.CombinedLoggingHandler(os.Stdout, router)
}
