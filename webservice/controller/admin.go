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

	// repo's
	courseRepo := &model.CourseRepository{}
	assignmentRepo := model.AssignmentRepository{}

	//get courses to user/teacher
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//need custom struct to get the coursecode
	type ActiveAssignment struct {
		Assignment model.Assignment
		CourseCode string
	}

	var activeAssignments []ActiveAssignment

	for _, course := range courses { //iterate all courses
		assignments, err := assignmentRepo.GetAllFromCourse(course.ID) //get assignments from course
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		for _, assignment := range assignments { //go through all it's assignments again
			activeAssignments = append(activeAssignments, ActiveAssignment{Assignment: assignment, CourseCode: course.Code})
		}

	}

	v := view.New(r)
	v.Name = "admin/index"

	v.Vars["Courses"] = courses
	v.Vars["Assignments"] = activeAssignments

	v.Render(w)
}
