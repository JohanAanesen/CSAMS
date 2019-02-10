package util

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"log"
	"net/http"
)

func IsTeacher(r *http.Request) bool {
	//check if user is already logged in
	user := GetUserFromSession(r)

	//check that user is a teacher
	if !user.Teacher { //not a teacher or logged in
		return false
	}

	return IsLoggedIn(r)
}

func IsLoggedIn(r *http.Request) bool {
	//get user from session
	user := GetUserFromSession(r)

	//check that user is a teacher
	if !user.Authenticated { //not logged in
		return false
	}

	return true
}

func GetUserFromSession(r *http.Request) model.User {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Println(err)
		return model.User{Authenticated: false}
	}

	val := session.Values["user"]
	var user = model.User{}
	user, ok := val.(model.User)
	if !ok {
		return model.User{Authenticated: false}
	}
	return user
}

func SaveUserToSession(user model.User, w http.ResponseWriter, r *http.Request) bool {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Println(err)
		return false
	}

	session.Values["user"] = user

	err = session.Save(r, w) //save session changes

	if err != nil {
		//todo log this event
		log.Fatal(err)

		//redirect somewhere
		return false
	}

	return true
}
