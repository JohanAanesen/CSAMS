package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	_ "github.com/go-sql-driver/mysql" //database driver
	"log"
	"net/http"
)

// AdminGET handles GET-request at /admin
func AdminGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/index"

	assignmentRepo := model.AssignmentRepository{}
	assignments, err := assignmentRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//course repo
	courseRepo := &model.CourseRepository{}

	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	v.Vars["Courses"] = courses
	v.Vars["Assignments"] = assignments

	v.Render(w)
}
