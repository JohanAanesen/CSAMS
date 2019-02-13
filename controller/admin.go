package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"log"
	"net/http"
)

// AdminGET handles GET-request at /admin
func AdminGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/index"

	// TODO (Svein): Add data to the page (courses, assignments, etc)

	v.Render(w)
}

// AdminCourseGET handles GET-request at /admin/course
func AdminCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/index"

	// TODO (Svein): Add data to the page

	v.Render(w)
}

// AdminCreateCourseGET handles GET-request at /admin/course/create
func AdminCreateCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/create"

	// TODO (Svein): Add data to the page (courses, assignments, etc)

	v.Render(w)
}

// AdminCreateCoursePOST handles POST-request at /admin/course/create
// Inserts a new course to the database
func AdminCreateCoursePOST(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//check if user is already logged in
	user := session.GetUserFromSession(r)

	course := model.Course{
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	//insert into database
	rows, err := db.GetDB().Query("INSERT INTO course(coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?)",
		course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID)

	if err != nil {
		//todo log error
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	IndexGET(w, r) //success redirect to homepage
}

// AdminUpdateCourseGET handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/update"

	// TODO (Svein): Add data to the page (courses, assignments, etc)

	v.Render(w)
}

// AdminUpdateCoursePOST handles POST-request at /admin/course/update/{id}
func AdminUpdateCoursePOST(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}
}

// AdminAssignmentGET handles GET-request at /admin/assignment
func AdminAssignmentGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/index"

	// TODO (Svein): Add data to the page (courses, assignments, etc)

	v.Render(w)
}
