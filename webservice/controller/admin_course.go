package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"log"
	"net/http"
	"strconv"
	"time"
)

// AdminCourseGET handles GET-request at /admin/course
func AdminCourseGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/index"

	courses, err := model.GetCoursesToUser(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

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
	user := session.GetUserFromSession(r)

	course := model.Course{
		Hash:        xid.NewWithTime(time.Now()).String(),
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	// TODO (Svein): Move this to model, in a function
	//insert into database
	result, err := db.GetDB().Exec("INSERT INTO course(hash, coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?, ?)",
		course.Hash, course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID)

	// Log error
	if err != nil {
		//todo log error
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get course id
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Convert from int64 to int
	course.ID = int(id)

	// Log createCourse in the database and give error if something went wrong
	lodData := model.Log{UserID: user.ID, Activity: model.CreatedCourse, CourseID: course.ID}
	if !model.LogToDB(lodData) {
		log.Fatal("Could not save createCourse log to database! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Add user to course
	if !model.AddUserToCourse(user.ID, course.ID) {
		log.Println("Could not add user to course! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log joinedCourse in the db and give error if something went wrong
	lodData = model.Log{UserID: user.ID, Activity: model.JoinedCourse, CourseID: course.ID}
	if !model.LogToDB(lodData) {
		log.Fatal("Could not save createCourse log to database! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//IndexGET(w, r) //success redirect to homepage
	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage
}

// AdminUpdateCourseGET handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/update"

	v.Render(w)
}

// AdminUpdateCoursePOST handles POST-request at /admin/course/update/{id}
func AdminUpdateCoursePOST(w http.ResponseWriter, r *http.Request) {

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

	assignmentRepo := model.AssignmentRepository{}
	assignments, err := assignmentRepo.GetAllFromCourse(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	course := model.GetCourse(id)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/assignments"

	v.Vars["Course"] = course
	v.Vars["Assignments"] = assignments

	v.Render(w)
}