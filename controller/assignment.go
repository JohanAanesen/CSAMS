package controller

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//AssignmentGET serves assignment page to users
func AssignmentGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

//AssignmentAutoGET serves the auto validation page to user
func AssignmentAutoGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

//AssignmentPeerGET serves the peer review page to users
func AssignmentPeerGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

// AssignmentUploadGET serves the upload page
func AssignmentUploadGET(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	if id == "" {
		log.Println("Error: id can't be empty! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Convert id from string to int
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	assignmentRepo := model.AssignmentRepository{}

	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	formRepo := model.FormRepository{}

	form, err := formRepo.GetFromAssignmentID(assignment.ID)

	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	description := github_flavored_markdown.Markdown([]byte(assignment.Description))

	// Set values
	v := view.New(r)
	v.Vars["Assignment"] = assignment
	v.Vars["Description"] = template.HTML(description)
	v.Vars["Fields"] = form.Fields
	v.Name = "assignment/upload"
	v.Render(w)

}

// AssignmentUploadPOST servers the
func AssignmentUploadPOST(w http.ResponseWriter, r *http.Request) {

	repos := r.FormValue("url")
	if repos == "" {
		log.Println("Error: url can't be empty! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	AssignmentAutoGET(w, r)
}
