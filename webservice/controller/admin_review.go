package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
	"net/http"
)

// AdminReviewGET handles GET-requests @ /admin/review
func AdminReviewGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var reviewRepo = model.ReviewRepository{}

	reviews, err := reviewRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	v := view.New(r)
	v.Name = "admin/review/index"

	v.Vars["Reviews"] = reviews

	v.Render(w)
}

// AdminReviewCreateGET handles GET-requests @ /admin/review/create
func AdminReviewCreateGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/review/create"

	v.Render(w)
}

// AdminReviewCreatePOST handles POST-requests @ /admin/review/create
func AdminReviewCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Get data from the form
	data := r.FormValue("data")
	fmt.Println(data)
	// Declare Form-struct
	var form = model.Form{}
	// Unmarshal the JSON-string sent from the form
	err := json.Unmarshal([]byte(data), &form)
	if err != nil {
		log.Println(err)
		return
	}
	// Declare empty slice for error messages
	var errorMessages []string

	// Check form name
	if form.Name == "" {
		errorMessages = append(errorMessages, "Form name cannot be blank.")
	}

	// Check number of fields
	if len(form.Fields) == 0 {
		errorMessages = append(errorMessages, "Form needs to have at least 1 field.")
	}

	// Check if any error messages has been appended
	if len(errorMessages) != 0 {
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		v := view.New(r)
		v.Name = "admin/review/create"

		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)

		v.Render(w)

		return
	}

	// Declare an empty Repository for Submission
	var repo = model.ReviewRepository{}
	// Insert data to database
	err = repo.Insert(form)
	if err != nil {
		log.Println(err)
		return
	}

	// Redirect to /admin/submission
	http.Redirect(w, r, "/admin/review", http.StatusFound)
}

// AdminReviewUpdateGET handles GET-requests @ /admin/review/update/{id:[0-9]+}
func AdminReviewUpdateGET(w http.ResponseWriter, r *http.Request) {

}

// AdminReviewUpdatePOST handles POST-requests @ /admin/review/update
func AdminReviewUpdatePOST(w http.ResponseWriter, r *http.Request) {

}