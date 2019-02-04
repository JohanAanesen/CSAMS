package handlers

import (
	"html/template"
	"log"
	"net/http"
	"../../db"
)


type User struct {
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
	//check if user is already logged in

	if session.Values["username"] == ""{
		http.Redirect(w, r, "/login", http.StatusFound)
	}





	var test User

	err = db.DB.QueryRow("SELECT id, name FROM users where id = ?", 1).Scan(&test.ID, &test.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//send user to login if no valid login cookies exist

	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/test.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, test)
}
