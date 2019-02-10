package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"net/http"
)

//LogoutHandler logs out logged in users
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}


	if !util.IsLoggedIn(r){ //not logged in, can't logout
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	session.Values["user"] = model.User{} //replace user object with an empty one
	session.Options.MaxAge = -1 //expire cookie

	err = session.Save(r, w) //save 'empty' session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
