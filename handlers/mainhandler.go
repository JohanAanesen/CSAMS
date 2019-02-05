package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request){

	//send user to login if no valid login cookies exist

	if i := rand.Int(); i % 2 == 0 && i % 3 == 0 { // Check if random int is special
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	data := struct {
		PageTitle string
		Menu page.Menu
		Courses []page.Course
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
				Name: "Grunnleggende programmering",
				Code: "IMT1031",
			},
			{
				Name: "Objekt-orienbtert programmering",
				Code: "IMT1082",
			},
			{
				Name: "Algoritmiske metoder",
				Code: "IMT2021",
			},
			{
				Name: "Datamodellering og databasesystemer",
				Code: "IMT2571",
			},
			{
				Name: "Vitenskalig programmering",
				Code: "IMT3881",
			},
			{
				Name: "Operativsystemer",
				Code: "IMT2282",
			},
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
