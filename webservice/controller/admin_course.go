package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/mail"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"log"
	"net/http"
	"strconv"
	"time"
)

// AdminCourseGET handles GET-request at /admin/course
func AdminCourseGET(w http.ResponseWriter, r *http.Request) {
	currentUser := session.GetUserFromSession(r)

	// Services
	courseService := service.NewCourseService(db.GetDB())

	courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/index"

	v.Vars["Courses"] = courses

	v.Render(w)
}

// AdminCreateCourseGET handles GET-request at /admin/course/create
func AdminCreateCourseGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/create"

	v.Render(w)
}

// AdminCreateCoursePOST handles POST-request at /admin/course/create
// Inserts a new course to the database
func AdminCreateCoursePOST(w http.ResponseWriter, r *http.Request) {
	//check if user is already logged in
	currentUser := session.GetUserFromSession(r)

	course := model.Course{
		Hash:        xid.NewWithTime(time.Now()).String(),
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
		Teacher:     currentUser.ID,
	}

	// Services
	courseService := service.NewCourseService(db.GetDB())
	services := service.NewServices(db.GetDB())

	var err error

	// Insert new course
	course.ID, err = courseService.Insert(course)
	if err != nil {
		log.Println("course service insert", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log create course to db
	err = services.Logs.InsertAdminCourse(currentUser.ID, course.ID)
	if err != nil {
		log.Println("log, create course ", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Add user to course
	err = courseService.AddUser(currentUser.ID, course.ID)
	if err == service.ErrUserAlreadyInCourse {
		http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
		return
	}

	if err != nil {
		log.Println("Could not add user to course! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log join course to db
	err = services.Logs.InsertJoinCourse(currentUser.ID, course.ID)
	if err != nil {
		log.Println("log, admin join course ", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//IndexGET(w, r) //success redirect to homepage
	http.Redirect(w, r, "/admin/course", http.StatusFound) //success redirect to homepage
}

// AdminUpdateCourseGET handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseGET(w http.ResponseWriter, r *http.Request) {
	// Get url variables
	vars := mux.Vars(r)
	// Convert string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("string convert id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	courseService := service.NewCourseService(db.GetDB())

	// Fetch course
	course, err := courseService.Fetch(id)
	if err != nil {
		log.Println("course service fetch", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Set content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
	v := view.New(r)
	v.Name = "admin/course/update"
	// View variables
	v.Vars["Course"] = course //attach course to template
	// Render view
	v.Render(w)
}

// AdminUpdateCoursePOST handles POST-request at /admin/course/update
func AdminUpdateCoursePOST(w http.ResponseWriter, r *http.Request) {
	// Get id from form and convert to integer
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("string convert id", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//get new variables from request
	newName := r.FormValue("name")
	newCode := r.FormValue("code")
	newDescription := r.FormValue("description")
	newSemester := r.FormValue("semester")

	//make sure they are not empty
	if newName == "" || newCode == "" || newSemester == "" {
		// TODO (Svein): Display error messages and the form.
		log.Println("some new data is empty, course update")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Services
	courseService := service.NewCourseService(db.GetDB())
	services := service.NewServices(db.GetDB())

	// Get current user
	currentUser := session.GetUserFromSession(r)

	//get course from database
	course, err := courseService.Fetch(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//update variables
	course.ID = id
	course.Name = newName
	course.Code = newCode
	course.Description = newDescription
	course.Semester = newSemester

	//save to database
	err = courseService.Update(*course)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log update course to db
	err = services.Logs.InsertAdminUpdateCourse(currentUser.ID, course.ID)
	if err != nil {
		log.Println("log, update course ", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/course", http.StatusFound)
}

// AdminCourseAllAssignments handles GET-request @ /course/{id:[0-9]+}/assignments
func AdminCourseAllAssignments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	assignmentService := service.NewAssignmentService(db.GetDB())
	courseService := service.NewCourseService(db.GetDB())

	// Fetch all assignments from course
	assignments, err := assignmentService.FetchFromCourse(id)
	if err != nil {
		log.Println("assignment service fetch from course", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//get course from database
	course, err := courseService.Fetch(id)
	if err != nil {
		log.Println("course service fetch", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/assignments"

	v.Vars["Course"] = course
	v.Vars["Assignments"] = assignments

	v.Render(w)
}

// AdminEmailCourseGET handles GET-request @ /course/email/{id:[0-9]+}
func AdminEmailCourseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	userService := service.NewUserService(db.GetDB())
	courseService := service.NewCourseService(db.GetDB())

	// Check if user is an participant of said class or a teacher
	inCourse, err := courseService.UserInCourse(session.GetUserFromSession(r).ID, id)
	if err != nil {
		log.Println("course service, user in course", err)
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Check if teacher is in course
	if !inCourse {
		log.Println("user not participant of class")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Get course from database
	course, err := courseService.Fetch(id)
	if err != nil {
		log.Println("course service fetch", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get all students from course
	users, err := userService.FetchAllFromCourse(id)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Declare variables
	var noOfPrivate int
	var noOfStudent int

	// Count number of private and student emails
	for _, user := range users {

		// Don't count with the teachers
		if !user.Teacher {

			// Count private if user have private email
			if user.EmailPrivate.Valid {
				noOfPrivate++
			} else {
				noOfStudent++
			}
		}
	}

	// Get all emails from students in course
	emails, err := userService.FetchAllStudentEmails(id)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/sendemail"

	v.Vars["Course"] = course
	v.Vars["Users"] = users
	v.Vars["NoOfEmails"] = len(emails)
	v.Vars["NoOfPrivate"] = noOfPrivate
	v.Vars["NoOfStudent"] = noOfStudent

	v.Render(w)
}

// AdminEmailCoursePOST handles POST-request @ /course/email/{id:[0-9]+}
func AdminEmailCoursePOST(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	subject := r.FormValue("subject")
	message := r.FormValue("message")

	if subject == "" || message == "" {
		log.Println("fields can't be empty")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	currentUser := session.GetUserFromSession(r)

	// Services
	userService := service.NewUserService(db.GetDB())
	loggingService := service.NewLogsService(db.GetDB())

	// Structs
	mailservice := mail.Mail{}

	// Get all emails from students in course
	emails, err := userService.FetchAllStudentEmails(courseID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Send mail to multiple recipients
	err = mailservice.SendMultipleRecipient(emails, subject, message)
	if err != nil {
		log.Println("sendmultiplerecipients", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log event to db
	err = loggingService.InsertEmailStudents(currentUser.ID, courseID, emails)
	if err != nil {
		log.Println("logging email to students", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	AdminEmailCourseGET(w, r)

}
