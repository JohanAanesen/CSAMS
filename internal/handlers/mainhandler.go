package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
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

	session, err := db.CookieStore.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	//check if user is logged in

	if getUser(session).Authenticated == false { //redirect to /login if not logged in
		//send user to login if no valid login cookies exist
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	/////////////////test///////////////////

	var test Test

	err = db.DB.QueryRow("SELECT id, name FROM users where id = ?", 1).Scan(&test.ID, &test.Name)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	data := struct {
		PageTitle   string
		Menu        page.Menu
		Courses     []page.Course
		LoadFormCSS bool
	}{
		PageTitle: "Homepage",

		Menu: util.LoadMenuConfig("configs/menu/site.json"),

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
		LoadFormCSS: false,
	}

	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}
