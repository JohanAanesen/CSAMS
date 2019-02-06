package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"html/template"
	"log"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request){
	//check that user is logged in and is admin/teacher

	//find classes admin/teacher own

	// Data for displaying on screen
	data := struct {
		PageTitle string
		Menu page.Menu
		Courses page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Homepage",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),
		Courses: util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	w.WriteHeader(http.StatusOK)

	//parse templates
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}

func AdminCourseHandler(w http.ResponseWriter, r *http.Request) {
	// Data for displaying on screen
	data := struct {
		PageTitle string
		Menu page.Menu
		Courses page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Dashboard - Courses",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),
		Courses: util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}

func AdminCreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/create.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu page.Menu
	}{
		PageTitle: "Dashboard - Create course",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),
	}); err != nil {
		log.Fatal(err)
	}
}

func AdminCreateCourseRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: talk to database and stuff
}

func AdminUpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/update.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu page.Menu
	}{
		PageTitle: "Dashboard - Update course",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),
	}); err != nil {
		log.Fatal(err)
	}
}

func AdminAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/assignment/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu page.Menu
	}{
		PageTitle: "Dashboard - Assignments",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),
	}); err != nil {
		log.Fatal(err)
	}
}