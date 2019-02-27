package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
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

	courses, err := model.GetCoursesToUser(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Set values
	v := view.New(r)
	v.Name = "index"

	v.Vars["Courses"] = courses
	v.Vars["Message"] = joinedCourse

	v.Render(w)
}

// JoinCoursePOST adds user to course
func JoinCoursePOST(w http.ResponseWriter, r *http.Request) {
	joinedCourse = ""

	// Check if course exists
	course := model.CourseExists(r.FormValue("courseID"))

	// If course ID == "", it doesn't exist
	if course.ID == -1 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Get user
	user := session.GetUserFromSession(r)

	// Check that user isn't in this class
	if model.UserExistsInCourse(user.ID, course.ID) {
		//joinedCourse = ""
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Add user to course if possible
	if !model.AddUserToCourse(user.ID, course.ID) {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Log joinedCourse in the database and give error if something went wrong
	lodData := model.Log{UserID: user.ID, Activity: model.JoinedCourse, CourseID: course.ID}
	if !model.LogToDB(lodData) {
		log.Fatal("Could not save JoinCourse log to database! (index.go)")
	}

	// Give feedback to user
	joinedCourse = course.Code + " - " + course.Name

	//IndexGET(w, r)
	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage
}
