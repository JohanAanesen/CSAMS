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

// AdminSubmissionGET handles GET-request to /admin/submission
func AdminSubmissionGET(w http.ResponseWriter, r *http.Request) {
	// Get all submissions from database
	subRepo := model.SubmissionRepository{}
	submissions, err := subRepo.GetAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header code and content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/submission/index"
	// Set view variables
	v.Vars["Submissions"] = submissions
	// Render view
	v.Render(w)
}

// AdminSubmissionCreateGET handles GET-request to /admin/submission/create
func AdminSubmissionCreateGET(w http.ResponseWriter, r *http.Request) {
	// Set header code and content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/submission/create"
	// Render view
	v.Render(w)
}

// AdminSubmissionCreatePOST handles POST-request to /admin/submission/create
func AdminSubmissionCreatePOST(w http.ResponseWriter, r *http.Request) {
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
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Set header content type and status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// Create new view
		v := view.New(r)
		// Set template file
		v.Name = "admin/submission/create"
		// Set view variables
		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)
		// Render view
		v.Render(w)
		return
	}

	// Declare an empty Repository for Submission
	var repo = model.SubmissionRepository{}
	// Insert data to database
	err = repo.Insert(form)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Redirect to /admin/submission
	http.Redirect(w, r, "/admin/submission", http.StatusFound)
}

// AdminSubmissionUpdateGET handles GET-request @ /admin/submission/update/{id:[0-9]+}
func AdminSubmissionUpdateGET(w http.ResponseWriter, r *http.Request) {
	// Get variables from request
	vars := mux.Vars(r)
	// Convert id string to int, check for errors
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form from database, based on the id
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

	// Set header content-type and code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create a view
	v := view.New(r)
	// Set template-file
	v.Name = "admin/submission/update"
	// View variables
	v.Vars["formJSON"] = string(formBytes)
	// Render view
	v.Render(w)
}

// AdminSubmissionUpdatePOST handles POST-request @ /admin/submission/update
func AdminSubmissionUpdatePOST(w http.ResponseWriter, r *http.Request) {
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
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Set header content-type and code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Create a view
		v := view.New(r)
		// Set template file
		v.Name = "admin/submission/update"
		// Set view variables
		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)
		// Render view
		v.Render(w)
		return
	}

	// Declare an empty Repository for Submission
	var repo = model.SubmissionRepository{}
	// Insert data to database
	err = repo.Update(form)
	if err != nil {
		log.Println("update submission", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Redirect to /admin/submission
	http.Redirect(w, r, "/admin/submission", http.StatusFound)
}

// AdminSubmissionDELETE handles DELETE-request @ /admin/submission/delete
func AdminSubmissionDELETE(w http.ResponseWriter, r *http.Request) {
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

	// Delete the submission from database, if error, set error messages, if ok, set success message
	repo := model.SubmissionRepository{}
	err = repo.Delete(temp.ID)
	if err != nil {
		msg.Code = http.StatusInternalServerError
		msg.Message = err.Error()
		msg.Location = "/admin/submission"
	} else {
		msg.Code = http.StatusOK
		msg.Message = "Deletion successful"
		msg.Location = "/admin/submission"
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
