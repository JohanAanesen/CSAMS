package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// AdminReviewGET handles GET-requests @ /admin/review
func AdminReviewGET(w http.ResponseWriter, r *http.Request) {
	// Set header content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	reviewService := service.NewReviewService(db.GetDB())
	reviews, err := reviewService.FetchAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
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
		// Convert form to JSON, handle error if occurs
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
		v.Name = "admin/review/create"
		// Set view variables
		v.Vars["Errors"] = errorMessages
		v.Vars["formJSON"] = string(formBytes)
		// Render view
		v.Render(w)
		return
	}

	// Create new review service
	reviewService := service.NewReviewService(db.GetDB())

	// Insert data to database
	_, err = reviewService.Insert(form)
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

	used, err := services.Review.IsUsed(form.ID)
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
		v.Name = "admin/review/update_used"

		// Set view variables
		v.Vars["Form"] = form

		// Render view
		v.Render(w)

		return
	}

	// Convert form to JSON
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
		// Convert form to JSON
		formBytes, err := json.Marshal(&form)
		if err != nil {
			log.Println("marshal form", err)
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

	reviewService := service.NewReviewService(db.GetDB())
	err = reviewService.Update(form)
	if err != nil {
		log.Println("update review form", err)
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
	reviewService := service.NewReviewService(db.GetDB())
	err = reviewService.Delete(temp.ID)
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

// AdminReviewUpdateWeightsGET func
func AdminReviewUpdateWeightsGET(w http.ResponseWriter, r *http.Request) {
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
	v.Name = "admin/review/update_weights"

	v.Vars["Form"] = form

	v.Render(w)
}

// AdminReviewUpdateWeightsPOST func
func AdminReviewUpdateWeightsPOST(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/admin/review", http.StatusFound)
}

// AdminReviewUpdateUsedPOST func
func AdminReviewUpdateUsedPOST(w http.ResponseWriter, r *http.Request) {
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

		if field.Type == "radio" {
			choices := r.FormValue(fmt.Sprintf("choices_%s", field.Name))
			field.Choices = strings.Split(choices, "\n")
		}

		err = services.Field.Update(field.ID, field)
		if err != nil {
			log.Println("services, field, update", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Redirect
	http.Redirect(w, r, "/admin/review", http.StatusFound)
}