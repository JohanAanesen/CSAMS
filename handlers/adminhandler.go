package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request){
	//check that user is logged in and is admin/teacher

	//find classes admin/teacher own

	data := PageData{
		PageTitle: "Homepage",
		Courses: []Course{
			{
				Name: "IMT1001",
				Assignments: []Assignment{
					{
						Name: "Assignment 1",
						Description: "Lorem ipsum dolor sit amet",
						Deadline: "20190516T120000Z",
					},
					{
						Name: "Assignment 2",
						Description: "Lorem ipsum dolor sit amet",
						Deadline: "20190516T120000Z",
					},
					{
						Name: "Assignment 3",
						Description: "Lorem ipsum dolor sit amet",
						Deadline: "20190516T120000Z",
					},
				},
			},
			{
				Name: "IMT2014",
				Assignments: []Assignment{
					{
						Name: "Assignment 1",
						Description: "Lorem ipsum dolor sit amet",
						Deadline: "20190516T120000Z",
					},
					{
						Name: "Assignment 2",
						Description: "Lorem ipsum dolor sit amet",
						Deadline: "20190516T120000Z",
					},
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
