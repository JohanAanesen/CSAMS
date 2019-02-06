package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	//check that user is logged in
	// code

	//fetch users information from server

	// TODO remove this with actual information
	type class struct {
		Code string
		Name string
		Link string
	}

	type user struct {
		Name           string
		PrimaryEmail   string
		SecondaryEmail string
		Classes        []class
		NoOfClasses    int
	}

	classes := []class{
		{"IMT1337", "Mobile Development", "#"},
		{"IMT4200", "Application Development", "#"},
		{"IMT8008", "Cloud Technologies", "#"},
		{"IMT0880", "WWW technology", "#"},
	}

	// I'm a bit unsure if this is the best solution, but it works for now
	var i = 0
	for i = range classes {
		i++
	}

	data := user{"Ola Nordmann", "olano@stud.ntnu.no", "olameister@gmail.com", classes, i}
	// TODO : end here

	//parse information with template
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/user.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, data)

}

func UserUpdateRequest(w http.ResponseWriter, r *http.Request) {

	// TODO BUG : the form is sending unvalidated input, this should not happen >:(

	// TODO : get user information and compare
	// TODO : actually change information
	// TODO : return errors

	// TODO : remove this
	uName := "Ola Nordmann"
	uSemail := "olameister@gmail.com"
	uPass := "123abc"

	name := r.FormValue("usersName")
	secondaryEmail := r.FormValue("secondaryEmail")
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	repeatPass := r.FormValue("repeatPass")

	// Name can not be changed to blank!
	if name == "" {
		// TODO : Return with error
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	} else if name == uName {
		// Do nothing, name is not changed
	} else {
		// TODO : Change name
	}

	// If email is empty, the user doesn't want to change it
	if secondaryEmail == "" {
		// Do nothing
	} else if secondaryEmail == uSemail {
		// Do nothing, nothing is changed
	} else {
		// TODO : Change secondary email
	}

	// If it's empty, the user doesn't want to change it
	if oldPass == "" {
		// Do nothing
	} else if oldPass != uPass {
		// TODO : Return with error, not matching the old password
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	} else if newPass != "" && repeatPass != "" && newPass == repeatPass && newPass != oldPass {
		// TODO : Change password
	}
}
