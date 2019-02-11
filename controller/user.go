package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"html/template"
	"log"
	"net/http"
)

type pageData = struct {
	User        model.User
	PageTitle   string
	Menu        page.Menu
	Navbar      page.Menu
	Courses     page.Courses
	NoOfClasses int
}

// UserGET serves user page to users
func UserGET(w http.ResponseWriter, r *http.Request) {

	user := session.GetUserFromSession(r)

	if !session.IsLoggedIn(r) {
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	courses := db.GetCoursesToUser(user.ID)

	// Data for displaying in view
	data := pageData{
		User:        user,
		PageTitle:   user.Name,
		Menu:        util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:      util.LoadMenuConfig("configs/menu/site.json"),
		Courses:     courses,
		NoOfClasses: len(courses.Items),
	}

	//parse templates
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/user.html")

	if err != nil {
		log.Println(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
	}
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
	hash := db.GetHash(user.ID)

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
	} else if name != user.Name && db.UpdateUserName(user.ID, name) {
		log.Println("Success: Name changed from " + user.Name + " to " + name)

		//update session
		user.Name = name
		session.SaveUserToSession(user, w, r)

		// Log name change in the database
		db.LogToDB(user.ID, db.ChangeName)
	}

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != user.EmailPrivate {
		if db.UpdateUserEmail(user.ID, secondaryEmail) {
			log.Println("Success: Private email changed from " + user.EmailPrivate + " to " + secondaryEmail)

			//update session
			user.EmailPrivate = secondaryEmail
			session.SaveUserToSession(user, w, r)

			// Log email change in the database
			db.LogToDB(user.ID, db.ChangeName)
		}
	}

	// No password input fields can be empty,
	// the new password has to be equal to repeat password field,
	// and the new password can't be the same as the old password
	passwordIsOkay := oldPass != "" && newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass

	// If there's no problem with passwords and teh password is changed
	if passwordIsOkay && db.CheckPasswordHash(oldPass, hash) {
		if db.UpdateUserPassword(user.ID, newPass) {
			log.Println("Success: Password is now changed!")

			// Log password change in the database
			db.LogToDB(user.ID, db.ChangePassword)
		} else {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	UserGET(w, r)
}

