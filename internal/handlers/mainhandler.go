package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"

	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"html/template"
	"log"
	"net/http"
)

//Test struct, should be removed soon
type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//MainHandler serves homepage to users
func MainHandler(w http.ResponseWriter, r *http.Request) {

	//check if user is logged in
	if util.GetUserFromSession(r).Authenticated == false { //redirect to /login if not logged in
		//send user to login if no valid login cookies exist
		//http.Redirect(w, r, "/login", http.StatusFound)
		LoginHandler(w, r)
		return
	}

	data := struct {
		PageTitle   string
		Menu        page.Menu
		Courses     page.Courses
		LoadFormCSS bool
	}{
		PageTitle:   "Homepage",
		Menu:        util.LoadMenuConfig("configs/menu/site.json"),
		Courses:     util.LoadCoursesConfig("configs/dd.json"),
		LoadFormCSS: false,
	}

	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/index.html")

	if err != nil {
		log.Println(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
	}
}
