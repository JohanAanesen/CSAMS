package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
	"net/http"
)

// IndexGET serves homepage to authenticated users, send anonymous to login
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Current User
	currentUser := session.GetUserFromSession(r)

	if !currentUser.Authenticated {
		LoginGET(w, r)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	//get courses to user
	courses, err := services.Course.FetchAllForUserOrdered(currentUser.ID)
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
		assignments, err := services.Assignment.FetchFromCourse(course.ID) //get assignments from course
		if err != nil {
			log.Println("services assignment, fetch from course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// TODO time-norwegian
		timeNow := util.GetTimeInCorrectTimeZone()
		for _, assignment := range assignments { //go through all it's assignments again
			if timeNow.After(assignment.Publish) && timeNow.Before(assignment.Deadline) { //save all 'active' assignments
				activeAssignments = append(activeAssignments, ActiveAssignment{Assignment: *assignment, CourseCode: course.Code})
			}
		}

	}

	// Set values
	v := view.New(r)
	v.Name = "index"

	v.Vars["Courses"] = courses
	v.Vars["Assignments"] = activeAssignments
	v.Vars["Message"] = session.GetAndDeleteMessageFromSession(w, r)

	v.Render(w)
}

// JoinCoursePOST adds user to course
func JoinCoursePOST(w http.ResponseWriter, r *http.Request) {
	// Get user
	currentUser := session.GetUserFromSession(r)
	// Services
	services := service.NewServices(db.GetDB())

	hash := r.FormValue("courseID")

	// Check if course exists
	course := services.Course.Exists(hash)
	// If course ID == "", it doesn't exist
	if course.ID == -1 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	ok, err := services.Course.UserInCourse(currentUser.ID, course.ID)
	if err != nil {
		log.Println("services, course, user in course", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	// Check that user isn't in this class
	if !ok {
		//joinedCourse = ""
		log.Println("user not in course")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	err = services.Course.AddUser(currentUser.ID, course.ID)
	// Add user to course if possible
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// TODO (Svein): Make this to a function, eg.: LogUserJoinedCourse(userID, courseID)
	// Log joinedCourse in the database and give error if something went wrong
	logData := model.Log{UserID: currentUser.ID, Activity: model.JoinedCourse, CourseID: course.ID}
	err = model.LogToDB(logData)

	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Give feedback to user
	session.SaveMessageToSession("You joined "+course.Code+" - "+course.Name, w, r)

	IndexGET(w, r)
	//http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage
}
