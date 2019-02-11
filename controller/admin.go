package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"html/template"
	"log"
	"net/http"
)

// AdminGET handles GET-request at /admin
func AdminGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//find classes admin/teacher own

	// Data for displaying on screen
	data := struct {
		PageTitle   string
		Menu        page.Menu
		Navbar      page.Menu
		Courses     page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Homepage",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
		Courses:   util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	w.WriteHeader(http.StatusOK)

	//parse templates
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/index.html")

	if err != nil {
		log.Println(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
	}
}

// AdminCourseGET handles GET-request at /admin/course
func AdminCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Data for displaying on screen
	data := struct {
		PageTitle   string
		Menu        page.Menu
		Navbar      page.Menu
		Courses     page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Dashboard - Courses",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
		Courses:   util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/course/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}

// AdminCreateCourseGET handles GET-request at /admin/course/create
func AdminCreateCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/course/create.html")

	if err != nil {
		ErrorHandler(w, r, 404)
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Create course",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
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

	course := page.Course{
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	//insert into database
	rows, err := db.DB.Query("INSERT INTO course(coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?)",
		course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID)

	if err != nil {
		//todo log error
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	IndexGET(w, r) //success redirect to homepage
}

// AdminUpdateCourseGET handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseGET(w http.ResponseWriter, r *http.Request) {
	//check that user is a teacher
	if !session.IsTeacher(r) { //not a teacher, error 401
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/course/update.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Update course",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
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
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/assignment/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Assignments",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
}
