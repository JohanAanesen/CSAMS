package handlers

import (
	"../internal/page"
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

		Menu: page.Menu{
			Items: []page.MenuItem{
				{Name: "Courses", Href: "/"},
				{Name: "Profile", Href: "/user"},
				{Name: "Admin Dashboard", Href: "/admin"},
			},
		},

		Courses: []page.Course{
			{
				Name: "IMT1001",
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
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/home.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}
