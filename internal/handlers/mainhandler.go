package handlers

import (
	"html/template"
	"log"
	"net/http"
	"../../db"
)


type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MainHandler(w http.ResponseWriter, r *http.Request){

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
		panic(err.Error()) // proper error handling instead of panic in your app
	}



	temp, err := template.ParseFiles("web/test.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, test)
}
