package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
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

	//course repo
	courseRepo := &model.CourseRepository{}

	//get courses to user
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "user/profile"

	v.Vars["User"] = user
	v.Vars["Courses"] = courses

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
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != user.EmailPrivate {
		if model.UpdateUserEmail(user.ID, secondaryEmail) {

			// Save information to log struct
			logData := model.Log{UserID: user.ID, Activity: model.ChangeEmail, OldValue: user.EmailPrivate, NewValue: secondaryEmail}

			//update session
			user.EmailPrivate = secondaryEmail
			session.SaveUserToSession(user, w, r)

			// Log email change in the database and give error if something went wrong
			err := model.LogToDB(logData)
			if err != nil {
				log.Println(err.Error())
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}
	}

	// No password input fields can be empty,
	// the new password has to be equal to repeat password field,
	// and the new password can't be the same as the old password
	passwordIsOkay := oldPass != "" && newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass

	// If there's no problem with passwords and the password is changed
	if passwordIsOkay && model.CheckPasswordHash(oldPass, hash) {
		if model.UpdateUserPassword(user.ID, newPass) {

			// Save information to log struct
			logData := model.Log{UserID: user.ID, Activity: model.ChangePassword}

			// Log password change in the database and give error if something went wrong
			err := model.LogToDB(logData)
			if err != nil {
				log.Println(err.Error())
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		} else {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	//UserGET(w, r)
	http.Redirect(w, r, "/user", http.StatusFound) //success redirect to homepage
}
