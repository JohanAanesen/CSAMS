package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
	"net/http"
	"strconv"
)

func AdminChangePassGET(w http.ResponseWriter, r *http.Request) {

	//course repo
	courseRepo := &model.CourseRepository{}
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/changepassword/index"
	v.Vars["Courses"] = courses

	v.Render(w)
}

func AdminChangePassPOST(w http.ResponseWriter, r *http.Request) {

	formVal := r.FormValue("course_id")

	// If courseID is blank, give error
	if formVal == "" {
		log.Println("error: course_id is nil")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Convert courseID to int
	courseID, err := strconv.Atoi(formVal)
	if err != nil {
		log.Printf(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get all students from courseID
	students := model.GetUsersToCourse(courseID)
	if len(students) < 0 {
		log.Println("Error: could not get students from course! (admin_change_pass.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get courses
	courseRepo := &model.CourseRepository{}
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/changepassword/index"
	v.Vars["Courses"] = courses
	v.Vars["Students"] = students
	v.Vars["SelectedCourse"] = courseID

	v.Render(w)

}
