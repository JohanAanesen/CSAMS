package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"net/http"
)

// TODO : maybe remove/refactor global variable later :/
var joinedCourse = ""

// IndexGET serves homepage to authenticated users, send anonymous to login
func IndexGET(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromSession(r)

	if !user.Authenticated {
		LoginGET(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Set values
	v := view.New(r)
	v.Name = "index"
	v.Vars["Auth"] = user.Authenticated
	v.Vars["Courses"] = db.GetCoursesToUser(user.ID)
	v.Vars["Message"] = joinedCourse
	v.Render(w)
}

// JoinCoursePOST adds user to course
func JoinCoursePOST(w http.ResponseWriter, r *http.Request) {
	joinedCourse = ""

	// Check if course exists
	course := db.CourseExists(r.FormValue("courseID"))

	// If course ID == -1, it doesn't exist
	if course.ID == -1 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Get user
	user := session.GetUserFromSession(r)

	// Check that user isn't in this class
	if db.UserExistsInCourse(user.ID, course.ID) {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Add user to course if possible
	if !db.AddUserToCourse(user.ID, course.ID) {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Give feedback to user
	joinedCourse = course.Code + " - " + course.Name

	IndexGET(w, r)
}
