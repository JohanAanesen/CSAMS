package handlers

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
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

// UserHandler serves user page to users
func UserHandler(w http.ResponseWriter, r *http.Request) {

	user := util.GetUserFromSession(r)

	if !util.IsLoggedIn(r) {
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

// UserUpdateRequest changes the user information
func UserUpdateRequest(w http.ResponseWriter, r *http.Request) {

	user := util.GetUserFromSession(r)

	if !util.IsLoggedIn(r) {
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
		fmt.Println("Success: Name changed from " + user.Name + " to " + name)

		//update session
		user.Name = name
		util.SaveUserToSession(user, w, r)
	}

	// Users Email
	// If secondary-email input isn't blank it has changed
	if secondaryEmail != "" && secondaryEmail != user.EmailPrivate {
		if db.UpdateUserEmail(user.ID, secondaryEmail) {
			fmt.Println("Success: Private email changed from " + user.EmailPrivate + " to " + secondaryEmail)

			//update session
			user.EmailPrivate = secondaryEmail
			util.SaveUserToSession(user, w, r)
		}
	}

	// No password input fields can be empty,
	// the new password has to be equal to repeat password field,
	// and the new password can't be the same as the old password
	passwordIsOkay := oldPass != "" && newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass

	// If there's no problem with passwords and teh password is changed
	if passwordIsOkay && db.CheckPasswordHash(oldPass, hash) {
		if db.UpdateUserPassword(user.ID, newPass) {
			fmt.Println("Success: Password is now changed!")
		} else {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	UserHandler(w, r)
}
