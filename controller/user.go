package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"log"
	"net/http"
)

// UserGET serves user page to users
func UserGET(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromSession(r)

	if !session.IsLoggedIn(r) {
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	courses := model.GetCoursesToUser(user.ID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "user/profile"

	v.Vars["Auth"] = user.Authenticated
	v.Vars["User"] = user
	v.Vars["Courses"] = courses // TODO (Svein): Take a look in user/profile.tmpl how this is used, and why noOfClasses is gone (Hint: {{len .Courses.Items}})

	v.Render(w)
}

// UserUpdatePOST changes the user information
func UserUpdatePOST(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromSession(r)

	if !session.IsLoggedIn(r) {
		//not logged in
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Get hashed password
	hash := model.GetHash(user.ID)

	// Get new data from the form
	name := r.FormValue("usersName")
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	// Users name
	if name == "" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	} else if name != user.Name && model.UpdateUserName(user.ID, name) {
		log.Println("Success: name changed from " + user.Name + " to " + name)

		// Save information to log struct
		logData := model.Log{UserID: user.ID, Activity: model.ChangeName, OldValue: user.Name, NewValue: name}

		//update session
		user.Name = name
		session.SaveUserToSession(user, w, r)

		// Log name change in the database and give error if something went wrong
		if !model.LogToDB(logData) {
			log.Fatal("Could not save name log to database! (userhandler.go)")
		}
	}

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != user.EmailPrivate {
		if model.UpdateUserEmail(user.ID, secondaryEmail) {
			log.Println("Success: Private email changed from " + user.EmailPrivate + " to " + secondaryEmail)

			// Save information to log struct
			logData := model.Log{UserID: user.ID, Activity: model.ChangeEmail, OldValue: user.EmailPrivate, NewValue: secondaryEmail}

			//update session
			user.EmailPrivate = secondaryEmail
			session.SaveUserToSession(user, w, r)

			// Log email change in the database and give error if something went wrong
			if !model.LogToDB(logData) {
				log.Fatal("Could not save email log to database! (userhandler.go)")
			}
		}
	}

	// No password input fields can be empty,
	// the new password has to be equal to repeat password field,
	// and the new password can't be the same as the old password
	passwordIsOkay := oldPass != "" && newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass

	// If there's no problem with passwords and teh password is changed
	if passwordIsOkay && model.CheckPasswordHash(oldPass, hash) {
		if model.UpdateUserPassword(user.ID, newPass) {
			log.Println("Success: Password is now changed!")

			// Save information to log struct
			logData := model.Log{UserID: user.ID, Activity: model.ChangePassword}

			// Log password change in the database and give error if something went wrong
			if !model.LogToDB(logData) {
				log.Fatal("Could not save password log to database! (userhandler.go)")
			}
		} else {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	//UserGET(w, r)
	http.Redirect(w, r, "/user", http.StatusFound) //success redirect to homepage
}
