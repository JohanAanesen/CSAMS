package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	_ "github.com/go-sql-driver/mysql" //database driver
	"log"
	"net/http"
)

// AdminGET handles GET-request at /admin
func AdminGET(w http.ResponseWriter, r *http.Request) {
	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Services
	courseService := service.NewCourseService(db.GetDB())
	assignmentService := service.NewAssignmentService(db.GetDB())

	// Get courses to current user
	courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println("course service, fetch all for user ordered", err)
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
		// Fetch assignments from course
		assignments, err := assignmentService.FetchFromCourse(course.ID)
		if err != nil {
			log.Println("assignment service, fetch from course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		for _, assignment := range assignments { //go through all it's assignments again
			activeAssignments = append(activeAssignments, ActiveAssignment{Assignment: *assignment, CourseCode: course.Code})
		}

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/index"

	v.Vars["Courses"] = courses
	v.Vars["Assignments"] = activeAssignments

	v.Render(w)
}
