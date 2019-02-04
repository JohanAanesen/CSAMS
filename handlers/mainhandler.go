package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Course struct {
	Name string
}

type PageData struct {
	PageTitle string
	Courses []Course
}

func MainHandler(w http.ResponseWriter, r *http.Request){

	//send user to login if no valid login cookies exist

	data := PageData{
		PageTitle: "Homepage",
		Courses: []Course{
			{"IMT1001"},
			{"IMT2001"},
			{"IMT3001"},
		},
	}

	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/home.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}
