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

	data := struct {
		PageTitle string
		Menu page.Menu
		Courses []page.Course
		Assignments []page.Assignment
	}{
		PageTitle: "Homepage",
		Menu: util.LoadMenuConfig("configs/menu/dashboard.json"),

		Courses: []page.Course{
			{
				Code: "IMT1001",
				Name: "Programming for dummies",
				Assignments: []page.Assignment{
					{Name: "Task 1", Description: "Code a \"Hello World program in Go\"", Deadline: "06/02/2019"},
					{Name: "Task 2", Description: "Code a Fibonacci-sequence both in a loop and a recursive functions", Deadline: "12/02/2019"},
					{Name: "Task 3", Description: "Make a Facebook-clone in Go", Deadline: "16/05/2019"},
				},
			},
		},
	}

	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}

func AdminCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", nil); err != nil {
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

	if err = temp.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Fatal(err)
	}
}

func AdminUpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/update.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", nil); err != nil {
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

	if err = temp.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Fatal(err)
	}
}