package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	_ "github.com/go-sql-driver/mysql" //database driver
	"github.com/rs/xid"
	"log"
	"net/http"
	"strconv"
	"time"
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

	assignmentRepo := model.AssignmentRepository{}
	assignments := assignmentRepo.GetAll()

	v.Vars["Assignments"] = assignments

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
		Hash:        xid.NewWithTime(time.Now()).String(),
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	// TODO (Svein): Move this to model, in a function
	//insert into database
	rows, err := db.GetDB().Query("INSERT INTO course(hash, coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?, ?)",
		course.Hash, course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID)

	if err != nil {
		//todo log error
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	/* TODO : get course id and add to logs
	// Log createCourse in the database and give error if something went wrong
	lodData := model.Log{UserID: user.ID, Activity: model.CreatedCourse, CourseID: idGoesHere}
	if !db.LogToDB(lodData) {
		log.Fatal("Could not save createCourse log to database! (admin.go)")
	}
	*/

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

// AdminAssignmentCreateGET handles GET-request from /admin/assigment/create
func AdminAssignmentCreateGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	var subRepo = model.SubmissionRepository{}
	submissions, err := subRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/create"

	v.Vars["Courses"] = []struct { // TODO (Svein): use database
		ID   int
		Code string
		Name string
	}{
		{ID: 1, Code: "IMT1031", Name: "Grunnleggende Programmering"},
		{ID: 2, Code: "IMT1082", Name: "Objekt-Orientert Programmering"},
		{ID: 3, Code: "IMT2021", Name: "Algoritmiske Metoder"},
	}

	v.Vars["Submissions"] = submissions

	v.Render(w)
}

// AdminAssignmentCreatePOST handles POST-request from /admin/assigment/create
func AdminAssignmentCreatePOST(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	assignmentRepository := model.AssignmentRepository{}

	publish, err := DatetimeLocalToRFC3339(r.FormValue("publish"))
	if err != nil {
		log.Println(err)
		return
	}

	deadline, err := DatetimeLocalToRFC3339(r.FormValue("deadline"))
	if err != nil {
		log.Println(err)
		return
	}

	courseIDString := r.FormValue("course_id")
	submissionIDString := r.FormValue("submission_id")
	reviewIDString := r.FormValue("review_id")

	courseID, err := strconv.Atoi(courseIDString)
	if err != nil {
		log.Println(err)
		return
	}

	var submissionID int
	if submissionIDString != "" {
		submissionID, err = strconv.Atoi(submissionIDString)
		if err != nil {
			log.Println(err)
			return
		}
	}

	var reviewID int
	if reviewIDString != "" {
		reviewID, err = strconv.Atoi(reviewIDString)
		if err != nil {
			log.Println(err)
			return
		}
	}

	assignment := model.Assignment{
		Name:         r.FormValue("name"),
		Description:  r.FormValue("description"),
		Publish:      publish,
		Deadline:     deadline,
		CourseID:     courseID,
		SubmissionID: submissionID,
		ReviewID:     reviewID,
	}

	success, err := assignmentRepository.Insert(assignment)
	if err != nil {
		log.Println(err)
		return
	}

	if success {
		// TODO (Svein): Celebrate
	}

}

// AdminSubmissionGET TODO (Svein): comment
func AdminSubmissionGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var subRepo = model.SubmissionRepository{}

	submissions, err := subRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	v := view.New(r)
	v.Name = "admin/submission/index"

	v.Vars["Submissions"] = submissions

	v.Render(w)
}

// AdminSubmissionCreateGET ...
func AdminSubmissionCreateGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/submission/create"

	// TODO (Svein): Add data to the page (courses, assignments, etc)

	v.Render(w)
}

// AdminSubmissionCreatePOST ...
func AdminSubmissionCreatePOST(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var form = model.Form{}

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
		return
	}

	var repo = model.SubmissionRepository{}

	err = repo.Insert(form)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO (Svein): Redirect or something!?
}

// DatetimeLocalToRFC3339 converts a string from datetime-local HTML input-field to time.Time object
func DatetimeLocalToRFC3339(str string) (time.Time, error) {
	// TODO (Svein): Move this to a utils.go or something
	if len(str) < 16 {
		return time.Time{}, errors.New("cannot convert a string less then 16 characters: DatetimeLocalToRFC3339()")
	}
	year := str[0:4]
	month := str[5:7]
	day := str[8:10]
	hour := str[11:13]
	min := str[14:16]

	value := fmt.Sprintf("%s-%s-%sT%s:%s:00Z", year, month, day, hour, min)
	return time.Parse(time.RFC3339, value)
}
