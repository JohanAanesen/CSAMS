package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
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
	currentUser := session.GetUserFromSession(r)

	// Services
	courseService := service.NewCourseService(db.GetDB())

	courses, err := courseService.FetchAllForUser(currentUser.ID)
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
	user := session.GetUserFromSession(r)

	course := model.Course{
		Hash:        xid.NewWithTime(time.Now()).String(),
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	// TODO (Svein): Move this to model, in a function/method
	//insert into database
	result, err := db.GetDB().Exec("INSERT INTO course(hash, coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?, ?)",
		course.Hash, course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID)

	// Log error
	if err != nil {
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
	logData := model.Log{UserID: user.ID, Activity: model.CreatedCourse, CourseID: course.ID}
	err = model.LogToDB(logData)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//course repo
	courseRepo := &model.CourseRepository{}

	// Add user to course
	if !courseRepo.AddUserToCourse(user.ID, course.ID) {
		log.Println("Could not add user to course! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log joinedCourse in the db and give error if something went wrong
	logData = model.Log{UserID: user.ID, Activity: model.JoinedCourse, CourseID: course.ID}
	err = model.LogToDB(logData)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//IndexGET(w, r) //success redirect to homepage
	http.Redirect(w, r, "/admin/course", http.StatusFound) //success redirect to homepage
}

// AdminUpdateCourseGET handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	//course repo
	courseRepo := &model.CourseRepository{}

	//get course from database
	course, err := courseRepo.GetSingle(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	v := view.New(r)
	v.Name = "admin/course/update"

	v.Vars["Course"] = course //attach course to template

	v.Render(w)
}

// AdminUpdateCoursePOST handles POST-request at /admin/course/update
func AdminUpdateCoursePOST(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//get new variables from request
	newName := r.FormValue("name")
	newCode := r.FormValue("code")
	newDescription := r.FormValue("description")
	newSemester := r.FormValue("semester")

	//make sure they are not empty
	if newName == "" || newCode == "" || newDescription == "" || newSemester == "" {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//course repo
	courseRepo := model.CourseRepository{}

	//get course from database
	course, err := courseRepo.GetSingle(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//update variables
	course.Name = newName
	course.Code = newCode
	course.Description = newDescription
	course.Semester = newSemester

	//save to database
	err = courseRepo.Update(id, course)
	if err != nil {
		log.Println(err)
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

	assignmentRepo := model.AssignmentRepository{}
	assignments, err := assignmentRepo.GetAllFromCourse(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//course repo
	courseRepo := model.CourseRepository{}

	//get course from database
	course, err := courseRepo.GetSingle(id)
	if err != nil {
		log.Println(err)
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
