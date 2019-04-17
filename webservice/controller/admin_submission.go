package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// AdminSubmissionGET handles GET-request to /admin/submission
func AdminSubmissionGET(w http.ResponseWriter, r *http.Request) {
	// Services
	submissionService := service.NewSubmissionService(db.GetDB())

	// Get all submissions from database
	submissions, err := submissionService.FetchAll()
	if err != nil {
		log.Println("get all submission", err)
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
		log.Println("unmarshal form from post request", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
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
			log.Println("json marshal form", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Set header content type and status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// Create view
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

	// Services
	submissionService := service.NewSubmissionService(db.GetDB())

	_, err = submissionService.Insert(form)
	if err != nil {
		log.Println("insert submission", err)
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
		log.Println("strconv atoi id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Get a single form based on ID from the database
	form, err := services.Form.Fetch(id)
	if err != nil {
		log.Println("form service get", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	used, err := services.Submission.IsUsed(form.ID)
	if err != nil {
		log.Println("services, review, is used", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if form is in use
	if used {
		// Set header content type and status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Create view
		v := view.New(r)
		// Set template file
		v.Name = "admin/submission/update_used"

		// Set view variables
		v.Vars["Form"] = form

		// Render view
		v.Render(w)
		return
	}

	// Convert form to JSON
	formBytes, err := json.Marshal(&form)
	if err != nil {
		log.Println("marshal form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header content-type and code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// Create view
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
		log.Println("unmarshal form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
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

		// Create view
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

	// Services
	submissionService := service.NewSubmissionService(db.GetDB())

	// Update form
	err = submissionService.Update(form)
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

	// Services
	services := service.NewServices(db.GetDB())

	// Delete the submission from database, if error, set error messages, if ok, set success message
	err = services.Submission.Delete(temp.ID)
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

// AdminSubmissionUpdateWeightsGET func
func AdminSubmissionUpdateWeightsGET(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, atoi", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	form, err := services.Form.Fetch(formID)
	if err != nil {
		log.Println("services, form, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header content-type and code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/submission/update_weights"

	v.Vars["Form"] = form

	v.Render(w)
}

// AdminSubmissionUpdateWeightsPOST func
func AdminSubmissionUpdateWeightsPOST(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, atoi", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	form, err := services.Form.Fetch(formID)
	if err != nil {
		log.Println("services, form, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("request, parse form", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	for _, field := range form.Fields {
		newWeight, err := strconv.Atoi(r.FormValue(field.Name))
		if err != nil {
			log.Println("strconv, atoi, request.FormValue(field.Name)", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		field.Weight = newWeight

		err = services.Field.Update(field.ID, field)
		if err != nil {
			log.Println("services, field, update", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Redirect
	http.Redirect(w, r, "/admin/submission", http.StatusFound)
}

// AdminSubmissionUpdateUsedPOST func
func AdminSubmissionUpdateUsedPOST(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, atoi", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	form, err := services.Form.Fetch(formID)
	if err != nil {
		log.Println("services, form, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("request, parse form", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	for _, field := range form.Fields {
		field.Label = r.FormValue(fmt.Sprintf("label_%s", field.Name))
		field.Description = r.FormValue(fmt.Sprintf("description_%s", field.Name))
		field.Required = r.FormValue(fmt.Sprintf("required_%s", field.Name)) == "on"

		choices := r.FormValue(fmt.Sprintf("choices_%s", field.Name))
		field.Choices = strings.Split(choices, "\n")

		err = services.Field.Update(field.ID, field)
		if err != nil {
			log.Println("services, field, update", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Redirect
	http.Redirect(w, r, "/admin/submission", http.StatusFound)
}
