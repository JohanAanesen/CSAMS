package controller

import (
	"encoding/gob"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
)

func init() {
	//todo maybe move this?
	gob.Register(model.User{})
}

//LoginGET serves login page to users
func LoginGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())

	// Models
	course := model.Course{}

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {
		// Check if course exists based on hash, and return it
		course = *courseService.Exists(hash)
		// Check if the hash is a valid hash
		if course.ID == -1 {
			log.Println("course id is -1")
			ErrorHandler(w, r, http.StatusBadRequest)
			hash = ""
			return
		}
	}

	// Check if user is already logged in
	currentUser := session.GetUserFromSession(r)
	if currentUser.Authenticated { //already logged in, redirect to homepage
		// If hash was valid, add user isn't in the course, then add user to course
		userInCourse, err := courseService.UserInCourse(currentUser.ID, course.ID)
		if err != nil {
			log.Println("course service, user in course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Check if user is not in course nad that hash is not blank
		if !userInCourse && hash != "" {
			// Add user to course
			err := courseService.AddUser(currentUser.ID, course.ID)
			if err == service.ErrUserAlreadyInCourse {
				http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
				return
			}

			// Check for errors
			if err != nil {
				log.Println("add user to course", err)
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}

		// Redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound) //redirect
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "login"

	// Send the correct link to template
	if hash == "" {
		v.Vars["Action"] = ""
	} else {
		v.Vars["Action"] = "?courseid=" + hash
	}

	v.Vars["Message"] = session.GetAndDeleteMessageFromSession(w, r)

	v.Render(w)
}

//LoginPOST validates login requests
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()
	// Get current user from session
	currentUser := session.GetUserFromSession(r)

	if currentUser.Authenticated { //already logged in, redirect to home page
		http.Redirect(w, r, "/", http.StatusFound) //todo redirect without 302
		return
	}

	// Services
	userService := service.NewUserService(db.GetDB())
	courseService := service.NewCourseService(db.GetDB())

	email := r.FormValue("email")       // email
	password := r.FormValue("password") // password
	hash := r.FormValue("courseid")     // courseID from link

	if email == "" || password == "" { //login credentials cannot be empty
		session.SaveMessageToSession("Credentials cannot be empty!", w, r)
		LoginGET(w, r)
		return
	}

	user, err := userService.Authenticate(p.Sanitize(email), p.Sanitize(password))
	if err != nil {
		log.Println("user service authenticate", err)
		session.SaveMessageToSession("Wrong credentials!", w, r)
		LoginGET(w, r) //try again
		return
	}

	//save user to session values
	user.Authenticated = true
	session.SaveUserToSession(user, w, r)

	// Add new user to course, if he's not in the course
	if hash != "" {
		course := courseService.Exists(hash)
		userInCourse, err := courseService.UserInCourse(currentUser.ID, course.ID)
		if err != nil {
			log.Println("course service, user in course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		if course.ID != -1 && !userInCourse {
			err := courseService.AddUser(user.ID, course.ID)

			if err == service.ErrUserAlreadyInCourse {
				http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
				return
			}

			if err != nil {
				log.Println("add user to course", err)
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}
	}

	// Removes any messages saved earlier, in failed attempts
	session.GetAndDeleteMessageFromSession(w, r)
	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage //todo change redirection
	//IndexGET(w, r) //redirect to homepage
}
