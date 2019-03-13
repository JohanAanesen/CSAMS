package controller

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// AdminReviewGET handles GET-requests @ /admin/review
func AdminReviewGET(w http.ResponseWriter, r *http.Request) {
	// Set header content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Get all reviews from database
	reviewRepo := model.ReviewRepository{}
	reviews, err := reviewRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/review/index"
	// Set view variables
	v.Vars["Reviews"] = reviews
	// Render view
	v.Render(w)
}

// AdminReviewCreateGET handles GET-requests @ /admin/review/create
func AdminReviewCreateGET(w http.ResponseWriter, r *http.Request) {
	// Set header content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/review/create"
	// Render view
	v.Render(w)
}

// AdminReviewCreatePOST handles POST-requests @ /admin/review/create
func AdminReviewCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Get data from the form
	data := r.FormValue("form_data")
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
		// Convert form to JSON, handle error if occurs
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Set header content type and status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// Create view
		v := view.New(r)
		// Set template file
		v.Name = "admin/review/create"
		// Set view variables
		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)
		// Render view
		v.Render(w)
		return
	}

	// Declare an empty Repository for Submission
	var repo = model.ReviewRepository{}
	// Insert data to database
	err = repo.Insert(form)
	if err != nil {
		log.Println("review insert", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Redirect to /admin/submission
	http.Redirect(w, r, "/admin/review", http.StatusFound)
}

// AdminReviewUpdateGET handles GET-requests @ /admin/review/update/{id:[0-9]+}
func AdminReviewUpdateGET(w http.ResponseWriter, r *http.Request) {
	// Get variables from the request
	vars := mux.Vars(r)
	// Convert id from string to id, and handle error
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// Get a single form based on ID from the database
	formRepo := model.FormRepository{}
	form, err := formRepo.Get(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// Convert form to JSON
	formBytes, err := json.Marshal(&form)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// Set header content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/review/update"
	// Set view variables
	v.Vars["formJSON"] = string(formBytes)
	// Render view
	v.Render(w)
}

// AdminReviewUpdatePOST handles POST-requests @ /admin/review/update
func AdminReviewUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Get data from the form
	data := r.FormValue("form_data")

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
		// Convert form to JSON
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Set header content type and status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// Create view
		v := view.New(r)
		// Set template-file
		v.Name = "admin/submission/update"
		// Set view variables
		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)
		// Render view
		v.Render(w)
		return
	}

	// Declare an empty Repository for Submission
	var repo = model.ReviewRepository{}
	// Insert data to database
	err = repo.Update(form)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Redirect to /admin/submission
	http.Redirect(w, r, "/admin/review", http.StatusFound)
}

// AdminReviewDELETE handles DELETE-request @ /admin/review/delete
func AdminReviewDELETE(w http.ResponseWriter, r *http.Request) {
	// Make a temporary struct for retrieving the json data
	temp := struct {
		ID int `json:"id"`
	}{}

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		log.Println("json decode error", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Make a temporary struct for the response body
	msg := struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		Location string `json:"location"`
	}{}

	// Delete the review from database, if error, set error messages, if ok, set success message
	repo := model.ReviewRepository{}
	err = repo.Delete(temp.ID)
	if err != nil {
		msg.Code = http.StatusInternalServerError
		msg.Message = err.Error()
		msg.Location = "/admin/review"
	} else {
		msg.Code = http.StatusOK
		msg.Message = "Deletion successful"
		msg.Location = "/admin/review"
	}

	// Write response code to header, and content type to JSON
	w.WriteHeader(msg.Code)
	w.Header().Set("Content-Type", "application/json")

	// Encode response
	err = json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Println("json encode error", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
