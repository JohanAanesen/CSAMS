package controller

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Combined holds answer and field
type Combined struct {
	Answer model.Answer
	Field  model.Field
}

// MergedAnswerField is a struct for merging the answer and field array
type MergedAnswerField struct {
	Items []Combined
}

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

// AssignmentSingleGET handles GET-request @ /assignment/{id:[0-9]+}
func AssignmentSingleGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignmentRepo := model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	descriptionMD := []byte(assignment.Description)
	description := github_flavored_markdown.Markdown(descriptionMD)

	delivered, err := assignmentRepo.HasUserSubmitted(assignmentID, session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasReview, err := assignmentRepo.HasReview(assignmentID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasAutoValidation, err := assignmentRepo.HasAutoValidation(assignmentID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	submissionReviews := model.GetReviewUserIDs(session.GetUserFromSession(r).ID, assignment.ID)

	//course repo
	courseRepo := &model.CourseRepository{}

	course, err := courseRepo.GetSingle(assignment.CourseID)
	if err != nil {
		log.Println("Something went wrong with getting course (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO time
	var isDeadlineOver = assignment.Deadline.Before(time.Now().UTC().Add(time.Hour))

	// TODO : make this dynamic
	var hasBeenValidated = true

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "assignment/index"

	v.Vars["Assignment"] = assignment
	v.Vars["Description"] = template.HTML(description)
	v.Vars["Delivered"] = delivered
	v.Vars["HasReview"] = hasReview
	v.Vars["HasAutoValidation"] = hasAutoValidation
	v.Vars["IsDeadlineOver"] = isDeadlineOver
	v.Vars["CourseID"] = course.ID
	v.Vars["Reviews"] = submissionReviews
	v.Vars["HasBeenValidated"] = hasBeenValidated

	v.Render(w)
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

	// Check for ID in url and give error if not
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

	// Get assignment and log possible error
	assignmentRepo := model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get form and log possible error
	formRepo := model.FormRepository{}
	form, err := formRepo.GetFromAssignmentID(assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//course repo
	courseRepo := &model.CourseRepository{}

	// Get course and log possible error
	course, err := courseRepo.GetSingle(assignment.CourseID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	answers, err := model.GetUserAnswers(session.GetUserFromSession(r).ID, assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	com := MergedAnswerField{}
	// Only merge if user has delivered
	if len(answers) > 0 {

		// Make sure answers and fields are same length before merging
		if len(answers) != len(form.Fields) {
			log.Println("Error: answers(" + strconv.Itoa(len(answers)) + ") is not equal length as fields(" + strconv.Itoa(len(form.Fields)) + ")! (assignment.go)")
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Merge field and answer if assignment is delivered

		for i := 0; i < len(form.Fields); i++ {
			com.Items = append(com.Items, Combined{
				Answer: answers[i],
				Field:  form.Fields[i],
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Set values
	v := view.New(r)
	v.Vars["Course"] = course
	v.Vars["Assignment"] = assignment
	v.Vars["Fields"] = form.Fields
	v.Vars["Delivered"] = len(answers)
	v.Vars["AnswersAndFields"] = com.Items
	v.Name = "assignment/upload"
	v.Render(w)

}

// AssignmentUploadPOST servers the
func AssignmentUploadPOST(w http.ResponseWriter, r *http.Request) {

	//XSS sanitizer
	p := bluemonday.UGCPolicy()

	// Check for ID in url and give error if not
	id := r.FormValue("id")
	if id == "" {
		log.Println("Error: id can't be empty! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Convert id from string to int
	assignmentID, err := strconv.Atoi(p.Sanitize(id))
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get assignment and log possible error
	assignmentRepo := model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if the deadline is reached TODO time
	var isDeadlineOver = assignment.Deadline.Before(time.Now().UTC().Add(time.Hour))
	if isDeadlineOver {
		log.Println("Error: Deadline is reached! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get form and log possible error
	formRepo := model.FormRepository{}
	form, err := formRepo.GetFromAssignmentID(assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that submission id is valid
	if !assignment.SubmissionID.Valid {
		log.Println("Error: Something went wrong with submissionID, its nil (assignment.go))")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if user has uploaded already or not
	delivered, err := assignmentRepo.HasUserSubmitted(assignmentID, session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Start to fill out user Submission struct
	userSub := model.UserSubmission{
		UserID:       session.GetUserFromSession(r).ID,
		SubmissionID: assignment.SubmissionID.Int64,
		AssignmentID: assignment.ID,
	}

	// Get answers WITH answerID if the user has delivered
	if delivered {
		userSub.Answers, err = model.GetUserAnswers(session.GetUserFromSession(r).ID, assignmentID)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Check that every form is filled an give error if not
	for index, field := range form.Fields {

		// Check if they are empty and give error if they are
		if r.FormValue(field.Name) == "" {
			log.Println("Error: assignment with form name '" + field.Name + "' can not be empty! (assignment.go)")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// If delivered, only change the value
		if delivered {
			sanitizedValue := p.Sanitize(r.FormValue(field.Name)) //sanitize input
			userSub.Answers[index].Value = sanitizedValue
		} else {
			// Else, create new answers array
			sanitizedValue := p.Sanitize(r.FormValue(field.Name)) //sanitize input

			userSub.Answers = append(userSub.Answers, model.Answer{
				Type:  field.Type,
				Value: sanitizedValue,
			})
		}

	}

	// Initiate error
	err = nil

	// Upload or update answers
	if !delivered {
		err = model.UploadUserSubmission(userSub)
	} else {
		err = model.UpdateUserSubmission(userSub)
	}

	// Check for error
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Serve front-end again
	AssignmentUploadGET(w, r)
}

// AssignmentUserSubmissionGET serves one suser submission to admin and the peer reviews
func AssignmentUserSubmissionGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get relevant user
	user := model.GetUser(userID)
	if user.Authenticated != true {
		log.Printf("Error: Could not get user (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get relevant assignment
	assignmentRepo := &model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Give error if user isn't teacher or reviewer for this user
	if !session.GetUserFromSession(r).Teacher && !model.UserIsReviewer(session.GetUserFromSession(r).ID, assignment.ID, assignment.SubmissionID.Int64, userID) {
		log.Println("Error: Unauthorized access!")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Get course and log possible error
	courseRepo := &model.CourseRepository{}
	course, err := courseRepo.GetSingle(assignment.CourseID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	formRepo := model.FormRepository{}
	form, err := formRepo.GetFromAssignmentID(assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	answers, err := model.GetUserAnswers(userID, assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	com := MergedAnswerField{}
	// Only merge if user has delivered
	if len(answers) > 0 {

		// Make sure answers and fields are same length before merging
		if len(answers) != len(form.Fields) {
			log.Println("Error: answers(" + strconv.Itoa(len(answers)) + ") is not equal length as fields(" + strconv.Itoa(len(form.Fields)) + ")! (assignment.go)")
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		// Merge field and answer if assignment is delivered

		for i := 0; i < len(form.Fields); i++ {
			com.Items = append(com.Items, Combined{
				Answer: answers[i],
				Field:  form.Fields[i],
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Set values
	v := view.New(r)
	v.Name = "assignment/submission"
	v.Vars["Assignment"] = assignment
	v.Vars["User"] = user
	v.Vars["Course"] = course
	v.Vars["Fields"] = form.Fields
	v.Vars["Delivered"] = len(answers)
	v.Vars["AnswersAndFields"] = com.Items
	v.Vars["IsTeacher"] = session.GetUserFromSession(r).Teacher
	v.Render(w)
}
