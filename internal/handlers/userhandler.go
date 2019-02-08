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

type payload struct {
	User    model.User
	Courses page.Courses
}

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

	// Check if user is logged in and get user information
	ok, _, payload := checkUserStatus(w, r)

	if !ok {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Data for displaying in view
	data := pageData{
		User:        payload.User,
		PageTitle:   payload.User.Name,
		Menu:        util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:      util.LoadMenuConfig("configs/menu/site.json"),
		Courses:     payload.Courses,
		NoOfClasses: len(payload.Courses.Items),
	}

	w.WriteHeader(http.StatusOK)

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

	ok, userID, payload := checkUserStatus(w, r)

	// Only do all this if it was possible to get userdata from DB
	if ok {

		// Get hashed password
		hash := db.GetHash(userID)

		everythingOk := false
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
		} else if name != payload.User.Name && db.UpdateUserName(userID, name) {
			fmt.Println("Success: Name changed from " + payload.User.Name + " to " + name)
			everythingOk = true
		}

		// Users Email
		// If secondary-email input isn't blank and not changed, error
		if secondaryEmail != "" && secondaryEmail == payload.User.EmailPrivate {
			// Do nothing
			everythingOk = true
		} else if secondaryEmail != "" && db.UpdateUserEmail(userID, secondaryEmail) {
			fmt.Println("Success: Private email changed from " + payload.User.EmailPrivate + " to " + secondaryEmail)
			everythingOk = true
		} else {
			ErrorHandler(w, r, http.StatusBadRequest)
		}

		// Users password
		if oldPass == "" { // If it's empty, the user doesn't want to change it
			// Do nothing
			everythingOk = true
			// if the password isn't the same as before and repeat and new is equal but not equal to oldpass, change password
		} else if newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass && db.CheckPasswordHash(oldPass, hash) {
			if db.UpdateUserPassword(userID, newPass) {
				fmt.Println("Success: Password is now changed!")
				everythingOk = true
			}
		} else {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		if everythingOk {
			UserHandler(w, r)
			return
		}
	}
}

// Checks if the user is logged in and gets the correct information
func checkUserStatus(w http.ResponseWriter, r *http.Request) (bool, int, payload) {

	var data = payload{}
	//check that user is logged in
	session, err := db.CookieStore.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return false, -1, data
	}

	//check if user is logged in
	user := getUser(session)
	if user.Authenticated == false { //redirect to /login if not logged in
		//send user to login if no valid login cookies exist
		http.Redirect(w, r, "/login", http.StatusFound)
		return false, -1, data
	}

	// Get user
	user2 := db.GetUser(user.ID)

	// Get users courses
	courses := db.GetCoursesToUser(user.ID)

	data = payload{User: user2, Courses: courses}

	// Return data
	return true, user.ID, data
}
