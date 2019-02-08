package handlers

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/structs"
	"html/template"
	"log"
	"net/http"
)

type userProfile struct {
	Name           string
	PrimaryEmail   string
	SecondaryEmail string
	Courses        []structs.CourseDB
	NoOfClasses    int
}

// UserHandler serves user page to users
func UserHandler(w http.ResponseWriter, r *http.Request) {

	// Check if user is logged in and get user information
	ok, _, _, data := checkUserStatus(w, r)

	// If everything went ok, execute page
	if ok {

		//parse information with template
		temp, err := template.ParseFiles("web/user.html")
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		temp.Execute(w, data)
	}
}

// UserUpdateRequest changes the user information
func UserUpdateRequest(w http.ResponseWriter, r *http.Request) {

	ok, userID, hash, payload := checkUserStatus(w, r)

	// Only do all this if it's a POST and it was possible to get userdata from DB
	if r.Method == http.MethodPost && ok {

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
		} else if name != payload.Name && db.UpdateUserName(userID, name) {
			fmt.Println("Success: Name changed from " + payload.Name + " to " + name)
			everythingOk = true
		}

		// Users Email
		// If secondaryemail input isn't blank and not changed, error
		if secondaryEmail != "" && secondaryEmail == payload.SecondaryEmail {
			// Do nothing
			everythingOk = true
		} else if secondaryEmail != "" && db.UpdateUserEmail(userID, secondaryEmail) {
			fmt.Println("Success: Private email changed from " + payload.SecondaryEmail + " to " + secondaryEmail)
			everythingOk = true
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
func checkUserStatus(w http.ResponseWriter, r *http.Request) (bool, int, string, userProfile) {

	//check that user is logged in
	session, err := db.CookieStore.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return false, -1, "", userProfile{}
	}

	//check if user is logged in
	user := getUser(session)
	if user.Authenticated == false { //redirect to /login if not logged in
		//send user to login if no valid login cookies exist
		http.Redirect(w, r, "/login", http.StatusFound)
		return false, -1, "", userProfile{}
	}

	// Get user
	_, name, emailStudent, teacher, emailPrivate, hash := db.GetUser(user.ID)

	// Get users courses
	courses := db.GetCoursesToUser(user.ID)

	// Count the courses
	i := 0
	for i = range courses {
		i++
	}

	if teacher != -1 {
		return true, user.ID, hash, userProfile{name, emailStudent, emailPrivate, courses, i}
	}

	return false, -1, "", userProfile{}
}
