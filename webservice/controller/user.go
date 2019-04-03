package controller

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
)

// UserGET serves user page to users
func UserGET(w http.ResponseWriter, r *http.Request) {
	// Services
	services := service.NewServices(db.GetDB())

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Get courses to user
	courses, err := services.Course.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "user/profile"

	v.Vars["User"] = currentUser
	v.Vars["Courses"] = courses

	v.Render(w)
}

// UserUpdatePOST changes the user information
func UserUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()
	// Get current user
	currentUser := session.GetUserFromSession(r)
	// Services
	services := service.NewServices(db.GetDB())

	// Get hashed password
	hash, err := services.User.FetchHash(currentUser.ID)
	if err != nil {
		log.Println("services, user, fetch hash", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get new data from the form
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != currentUser.EmailPrivate.String {
		updatedUser := currentUser
		updatedUser.EmailPrivate = sql.NullString{
			String: p.Sanitize(secondaryEmail),
			Valid:  secondaryEmail != "",
		}

		// Check if the email exists in the db
		exist, _, err := services.User.EmailExists(p.Sanitize(secondaryEmail))
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// If the email already exists
		if exist {
			log.Println("email already exist in db")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		err = services.User.Update(currentUser.ID, updatedUser)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Save information to log struct
		logData := model.Log{UserID: currentUser.ID, Activity: model.ChangeEmail, OldValue: currentUser.EmailPrivate.String, NewValue: secondaryEmail}

		//update session
		currentUser.EmailPrivate = updatedUser.EmailPrivate
		session.SaveUserToSession(currentUser, w, r)

		// Log email change in the database and give error if something went wrong
		err = model.LogToDB(logData) // TODO (svein): Make this to a function, eg.: func LogChangeEmail(userID, oldValue, newValue) error {}
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// if no of these fields are empty, try to change password
	if oldPass != "" && newPass != "" && repeatPass != "" {

		// check if the old password is correct
		err = util.CompareHashAndPassword(oldPass, hash)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// If password hasn't changed or new and repeat doesn't match
		if newPass == oldPass || newPass != repeatPass {
			log.Println("New password has not changed or repeat password and new password is not the same")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Update password
		err = services.User.UpdatePassword(currentUser.ID, newPass)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Save information to log struct
		logData := model.Log{UserID: currentUser.ID, Activity: model.ChangePassword}

		// Log password change in the database and give error if something went wrong
		err = model.LogToDB(logData) // TODO (Svein): Make this logging into a function: func LogChangePassword(userID, oldHash?, newHash?) error {}
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

	}

	//UserGET(w, r)
	http.Redirect(w, r, "/user", http.StatusFound) //success redirect to homepage
}
