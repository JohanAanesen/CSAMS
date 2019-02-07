package handlers

import (
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
	ok, _,  _, data := checkUserStatus(w, r)

	// If everything went ok, execute page
	if ok {
		//parse information with template
		w.WriteHeader(http.StatusOK)

		temp, err := template.ParseFiles("web/user.html")
		if err != nil {
			log.Fatal(err)
		}

		temp.Execute(w, data)
	}
}

// UserUpdateRequest changes the user information
func UserUpdateRequest(w http.ResponseWriter, r *http.Request) {

	// TODO BUG : the form is sending unvalidated input, this should not happen >:(
	// TODO BUG : hash is empty in this function, but not the payload even if I add the hash to the payload :/

	// TODO : actually change information
	// TODO : return errors

	// Get new data from the form
	name := r.FormValue("usersName")
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	ok, userID, hash, payload := checkUserStatus(w, r)

	// Only do all this if it's a POST and it was possible to get userdata from DB
	if r.Method == http.MethodPost && ok {

		if name == "" { // Name can not be changed to blank!

			ErrorHandler(w, r, http.StatusBadRequest) // TODO : Return with error
			return
		} else if name == payload.Name {
			// Do nothing, name is not changed
		} else {
			db.UpdateUserName(userID, name)
		}

		if secondaryEmail == "" { // If email is empty, the user doesn't want to change it
			// Do nothing
		} else if secondaryEmail == payload.SecondaryEmail {
			// Do nothing, nothing is changed
		} else {
			db.UpdateUserEmail(userID, secondaryEmail)
		}

		if oldPass == "" { // If it's empty, the user doesn't want to change it
			// Do nothing
		} else if !db.CheckPasswordHash(oldPass, hash) { // Password is not the same as the current
			ErrorHandler(w, r, http.StatusBadRequest)

		} else if newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass {
			db.UpdateUserPassword(userID, newPass)
		} else {
			ErrorHandler(w, r, http.StatusBadRequest)
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

	_, name, emailStudent, teacher, emailPrivate, _ := db.GetUser(user.ID)
	hash2 := db.GetHash(user.ID)
	courses := db.GetCoursesToUser(user.ID)

	i := 0
	for i = range courses {
		i++
	}

	if teacher != -1 {
		return true, user.ID, hash2, userProfile{name, emailStudent, emailPrivate, courses, i}
	}

	return false, -1, "", userProfile{}
}

func updateUserName(name string) bool {

	return false
}

func updateEmailPrivate(email string) bool {

	return false
}

func updatePassword(password string) bool {

	return false
}
