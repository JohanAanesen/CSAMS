package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"log"
	"net/http"
)

//LogoutGET logs out logged in users
func LogoutGET(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Instance(r) //get session

	if err != nil {
		log.Println("get session error: ", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if !session.IsLoggedIn(r) { //not logged in, can't logout
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	//session.Empty(sess) // Empty all values in the cookie
	sess.Values["user"] = model.User{} //replace user object with an empty one
	sess.Options.MaxAge = -1           //expire cookie

	err = sess.Save(r, w)
	if err != nil {
		log.Println("save session error: ", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
