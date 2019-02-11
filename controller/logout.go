package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"net/http"
)

//LogoutGET logs out logged in users
func LogoutGET(w http.ResponseWriter, r *http.Request) {
	sess, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}

	if !session.IsLoggedIn(r) { //not logged in, can't logout
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	sess.Values["user"] = model.User{} //replace user object with an empty one
	sess.Options.MaxAge = -1           //expire cookie

	err = sess.Save(r, w) //save 'empty' session
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//todo log this event
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

