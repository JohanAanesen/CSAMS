package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"github.com/rs/xid"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Struct for keeping the frequently asked questions under /admin/faq
type faq struct {
	date      time.Time // Last edited time
	questions string    // The markdown with questions and answers
}

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
	if !db.LogToDB(lodData) {
		log.Fatal("Could not save createCourse log to database! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Add user to course
	if !db.AddUserToCourse(user.ID, course.ID) {
		log.Println("Could not add user to course! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log joinedCourse in the db and give error if something went wrong
	lodData = model.Log{UserID: user.ID, Activity: model.JoinedCourse, CourseID: course.ID}
	if !db.LogToDB(lodData) {
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

// AdminFaqGET handles GET-request at admin/faq/index
func AdminFaqGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	content := faq{questions: "-1"}

	// TODO : it feels wrong to have this here, but I think this is correct :S

	//insert into database
	rows, err := db.GetDB().Query("SELECT timestamp, questions FROM `adminfaq` WHERE id = 1") // OBS! Id is always 1 since it's only one entry in the table

	// Log error
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var timestamp time.Time
		var questions string

		rows.Scan(&timestamp, &questions)

		content = faq{
			date:      timestamp,
			questions: questions,
		}
	}

	if content.questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	questions := github_flavored_markdown.Markdown([]byte(content.questions))

	v := view.New(r)
	v.Name = "admin/faq/index"
	v.Vars["Updated"] = content.date.Format("02. January 2006")
	v.Vars["Questions"] = template.HTML(questions)

	v.Render(w)
}
