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
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
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
	assignments, err := assignmentRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	v.Vars["Assignments"] = assignments

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
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/course/update"

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

	assignmentRepo := model.AssignmentRepository{}

	assignments, err := assignmentRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/index"

	v.Vars["Assignments"] = assignments

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

	v.Vars["Courses"] = model.GetCoursesToUser(session.GetUserFromSession(r).ID)

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

	if publish.After(deadline) {
		// TODO (Svein): Give feedback on this. Also add this in the Javascript
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form values
	courseIDString := r.FormValue("course_id")
	submissionIDString := r.FormValue("submission_id")
	reviewIDString := r.FormValue("review_id")

	// String converted into integer
	courseID, err := strconv.Atoi(courseIDString)
	if err != nil {
		log.Println(err)
		return
	}

	var submissionID int
	// String converted into integer
	if submissionIDString != "" {
		submissionID, err = strconv.Atoi(submissionIDString)
		if err != nil {
			log.Println(err)
			return
		}
	}

	var reviewID int
	// String converted into integer
	if reviewIDString != "" {
		reviewID, err = strconv.Atoi(reviewIDString)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Put all data into an Assignment-struct
	assignment := model.Assignment{
		Name:         r.FormValue("name"),
		Description:  r.FormValue("description"),
		Publish:      publish,
		Deadline:     deadline,
		CourseID:     courseID,
		SubmissionID: submissionID,
		ReviewID:     reviewID,
	}

	// Insert data to database
	err = assignmentRepository.Insert(assignment)
	if err != nil {
		log.Println(err)
		return
	}

	// Redirect client to '/'
	http.Redirect(w, r, "/", http.StatusFound)
}

// AdminSubmissionGET handles GET-request to /admin/submission
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

// AdminSubmissionCreateGET handles GET-request to /admin/submission/create
func AdminSubmissionCreateGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/submission/create"

	v.Render(w)
}

// AdminSubmissionCreatePOST handles POST-request to /admin/submission/create
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

	http.Redirect(w, r, "/", http.StatusFound)
}

// DatetimeLocalToRFC3339 converts a string from datetime-local HTML input-field to time.Time object
func DatetimeLocalToRFC3339(str string) (time.Time, error) {
	// TODO (Svein): Move this to a utils.go or something
	if str == "" {
		return time.Time{}, errors.New("error: could not parse empty datetime-string")
	}
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

// AdminFaqGET handles GET-request at admin/faq/index
func AdminFaqGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Convert to html
	questions := github_flavored_markdown.Markdown([]byte(content.Questions))

	v := view.New(r)
	v.Name = "admin/faq/index"
	v.Vars["Updated"] = content.Date.Format("02. January 2006 - 15:04")
	v.Vars["Questions"] = template.HTML(questions)

	v.Render(w)
}

// AdminFaqEditGET returns the edit view for the faq
func AdminFaqEditGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/edit"
	v.Vars["Updated"] = content.Date.Format("02. January 2006 - 15:04")
	v.Vars["RawContent"] = content.Questions

	v.Render(w)
}

// AdminFaqUpdatePOST handles the edited markdown faq
func AdminFaqUpdatePOST(w http.ResponseWriter, r *http.Request) {

	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Check that the questions arrived
	updatedFAQ := r.FormValue("rawQuestions")
	if updatedFAQ == "" {
		log.Println("Form is empty! (admin.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check that it's possible to get the old faq from db
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that it's changes to the new faq
	if content.Questions == updatedFAQ {
		log.Println("Old and new faq can not be equal! (admin.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check that it went okay to add new faq to db
	if !model.UpdateFAQ(updatedFAQ) {
		log.Println("Something went wrong with updating the faq! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user for logging purposes
	user := session.GetUserFromSession(r)

	// Collect the log data
	logData := model.Log{
		UserID:   user.ID,
		Activity: model.UpdateAdminFAQ,
		OldValue: content.Questions,
		NewValue: updatedFAQ,
	}

	// Log that a teacher has changed the faq
	if !model.LogToDB(logData) {
		log.Println("Something went wrong with logging the new faq! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	AdminFaqGET(w, r)
}
