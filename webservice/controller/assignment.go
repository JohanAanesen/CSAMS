package controller

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

// AssignmentSingleGET handles GET-request @ /assignment/{id:[0-9]+}
func AssignmentSingleGET(w http.ResponseWriter, r *http.Request) {
	// Get current user
	currentUser := session.GetUserFromSession(r)
	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	delivered, err := services.SubmissionAnswer.HasUserSubmitted(assignmentID, currentUser.ID)
	if err != nil {
		log.Println("services, submission answer, has user submitted", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasReview, err := services.Assignment.HasReview(assignmentID)
	if err != nil {
		log.Println("services, assignment, has review", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasAutoValidation, err := services.Assignment.HasAutoValidation(assignmentID)
	if err != nil {
		log.Println("services, assignment, has auto validation", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Filter out the reviews that the current user already has done
	reviewUsers, err := services.Review.FetchReviewUsers(currentUser.ID, assignment.ID)
	if err != nil {
		log.Println("services, review, fetch review users", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	filteredSubmissionReviews := make([]model.User, 0)
	for _, user := range reviewUsers {
		check, err := services.ReviewAnswer.HasBeenReviewed(user.ID, currentUser.ID, assignmentID)
		if err != nil {
			log.Println("services, review answer, has been reviewed", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		if !check {
			filteredSubmissionReviews = append(filteredSubmissionReviews, *user)
		}
	}

	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("services, course, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO (Svein): Check if this is correct
	myReviews, err := services.ReviewAnswer.FetchForUser(currentUser.ID, assignment.ID)
	if err != nil {
		log.Println("services, review answer, fetch for reviewer", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO time-norwegian
	var isDeadlineOver = assignment.Deadline.Before(util.GetTimeInCorrectTimeZone())

	// TODO : make this dynamic
	var hasBeenValidated = false

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	v.Name = "assignment/index"

	v.Vars["Assignment"] = assignment
	v.Vars["Delivered"] = delivered
	v.Vars["HasReview"] = hasReview
	v.Vars["HasAutoValidation"] = hasAutoValidation
	v.Vars["IsDeadlineOver"] = isDeadlineOver
	v.Vars["CourseID"] = course.ID
	v.Vars["Reviews"] = filteredSubmissionReviews
	v.Vars["HasBeenValidated"] = hasBeenValidated
	v.Vars["MyReviews"] = myReviews // TODO (Svein): Fix this, all in one slice, split into N-slices based on reviewer
	v.Vars["IsTeacher"] = currentUser.Teacher

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
		log.Println("string convert atoi id", err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Get assignment and log possible error
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("get single assignment", err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	delivered, err := services.SubmissionAnswer.HasUserSubmitted(assignmentID, currentUser.ID)
	if err != nil {
		log.Println("services, submission answer, has user submitted", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get course and log possible error
	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("get single course", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	answers, err := services.SubmissionAnswer.FetchUserAnswers(currentUser.ID, assignment.ID)
	if err != nil {
		log.Println("get user answers", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Only merge if user has delivered
	if len(answers) > 0 {
		// Make sure answers and fields are same length before merging
		if len(answers) != len(submissionForm.Form.Fields) {
			log.Println("Error: answers(" + strconv.Itoa(len(answers)) + ") is not equal length as fields(" + strconv.Itoa(len(submissionForm.Form.Fields)) + ")! (assignment.go)")
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		for index, field := range submissionForm.Form.Fields {
			answers[index].HasComment = field.HasComment
			answers[index].Description = field.Description
		}
	} else {
		for _, item := range submissionForm.Form.Fields {
			answers = append(answers, &model.SubmissionAnswer{
				Type:        item.Type,
				Choices:     item.Choices,
				Weight:      item.Weight,
				Label:       item.Label,
				HasComment:  item.HasComment,
				Description: item.Description,
				Name:        item.Name,
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	v.Name = "assignment/upload"

	v.Vars["Course"] = course
	v.Vars["Assignment"] = assignment
	v.Vars["Fields"] = submissionForm.Form.Fields
	v.Vars["Delivered"] = delivered
	v.Vars["Answers"] = answers

	v.Render(w)
}

// AssignmentUploadPOST handles POST-request @ /assignment/submission/update
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
		log.Println("strconv atoi id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Services
	assignmentService := service.NewAssignmentService(db.GetDB())
	submissionAnswerService := service.NewSubmissionAnswerService(db.GetDB())
	submissionService := service.NewSubmissionService(db.GetDB())

	// Get assignment and log possible error
	assignment, err := assignmentService.Fetch(assignmentID)
	if err != nil {
		log.Println("assignment service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if the deadline is reached TODO fix this quick fix time-norwegian
	var isDeadlineOver = assignment.Deadline.Before(util.GetTimeInCorrectTimeZone())
	if isDeadlineOver {
		log.Println("Error: Deadline is reached! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check that submission id is valid
	if !assignment.SubmissionID.Valid {
		log.Println("Error: Something went wrong with submissionID, it is not valid (assignment.go))")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := submissionService.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("submission service, fetch from assignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if user has uploaded already or not
	delivered, err := submissionAnswerService.HasUserSubmitted(assignmentID, currentUser.ID)
	if err != nil {
		log.Println("submission answer service, has user submitted", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Create empty submission answer slice
	submissionAnswers := make([]*model.SubmissionAnswer, 0)

	// Get answers WITH answerID if the user has delivered
	if delivered {
		// Fetch answers from the user
		submissionAnswers, err = submissionAnswerService.FetchUserAnswers(currentUser.ID, assignment.ID)
		if err != nil {
			log.Println("submission answer service, fetch user answers", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Parse form from request
	err = r.ParseForm()
	if err != nil {
		log.Println("request parse form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that every form is filled an give error if not
	for index, field := range submissionForm.Form.Fields {
		// Check if they are empty and give error if they are
		if r.FormValue(field.Name) == "" && field.Type != "checkbox" && field.Type != "paragraph" && field.Type != "multi-checkbox" {
			log.Println("Error: assignment with form name '" + field.Name + "' can not be empty! (assignment.go)")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Initialize empty answer
		answer := model.SubmissionAnswer{}
		// Check if the field has comment enabled
		if field.HasComment {
			// Get comment content, sanitized
			answer.Comment.String = p.Sanitize(r.FormValue(field.Name + "_comment"))
			answer.Comment.Valid = answer.Comment.String != ""
		}

		// Check if multi-checkbox
		if field.Type == "multi-checkbox" {
			val := r.Form[field.Name]
			answer.Answer = p.Sanitize(strings.Join(val, ","))
		} else {
			// Sanitize input
			answer.Answer = p.Sanitize(r.FormValue(field.Name))
		}

		// Get field type
		answer.Type = field.Type
		answer.Name = field.Name
		answer.Label = field.Label
		answer.Description = field.Description
		answer.HasComment = field.HasComment

		// If delivered, only change the value
		if delivered {
			submissionAnswers[index].Answer = answer.Answer
			submissionAnswers[index].Comment = answer.Comment
			// Set name & label
			submissionAnswers[index].Name = field.Name
			submissionAnswers[index].Label = field.Label
			submissionAnswers[index].Description = field.Description
			submissionAnswers[index].HasComment = field.HasComment
		} else {
			// Else, create new answers array
			submissionAnswers = append(submissionAnswers, &answer)
		}

	}

	// Update user, assignment & submission id for all answers
	for _, item := range submissionAnswers {
		item.UserID = currentUser.ID
		item.AssignmentID = assignment.ID
		item.SubmissionID = int(assignment.SubmissionID.Int64)
	}

	// Upload or update answers
	if !delivered {
		err = submissionAnswerService.Upload(submissionAnswers)
	} else {
		err = submissionAnswerService.Update(submissionAnswers)
	}

	// Check for error
	if err != nil {
		log.Println("submission answer service, upload/update", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Serve front-end again
	AssignmentUploadGET(w, r)
}

// AssignmentUserSubmissionGET serves one user submission to admin and the peer reviews
func AssignmentUserSubmissionGET(w http.ResponseWriter, r *http.Request) {
	// Get parameters in the URL
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Printf("userid: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get relevant user
	user := model.GetUser(userID)
	if !user.Authenticated {
		log.Printf("Error: Could not get user (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Current user
	currentUser := session.GetUserFromSession(r)

	// Get relevant assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("get single assignment", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasBeenReviewed, err := services.ReviewAnswer.HasBeenReviewed(user.ID, currentUser.ID, assignment.ID)
	if err != nil {
		log.Println("has been reviewed", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if hasBeenReviewed && !currentUser.Teacher {
		IndexGET(w, r)
		return
	}

	// Give error if user isn't teacher or reviewer for this user
	isUserTheReviewer, err := services.Review.IsUserTheReviewer(currentUser.ID, userID, assignment.ID)
	if err != nil {
		log.Println("services, review, is user the reviewer", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//if !currentUser.Teacher && !model.UserIsReviewer(currentUser.ID, assignment.ID, assignment.SubmissionID.Int64, userID) {
	if !currentUser.Teacher && !isUserTheReviewer {
		log.Println("Error: Unauthorized access!")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// Get course and log possible error
	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("course get single", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	form, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission from form assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get assignmentAnswers to user if he has delivered
	assignmentAnswers, err := services.SubmissionAnswer.FetchUserAnswers(userID, assignment.ID)
	if err != nil {
		log.Println("get user assignmentAnswers", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	/*
		com := MergedAnswerField{}
		// Only merge if user has delivered
		if len(assignmentAnswers) > 0 {
			// Make sure assignmentAnswers and fields are same length before merging
			if len(assignmentAnswers) != len(form.Form.Fields) {
				log.Println("Error: assignmentAnswers(" + strconv.Itoa(len(assignmentAnswers)) + ") is not equal length as fields(" + strconv.Itoa(len(form.Form.Fields)) + ")! (assignment.go)")
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}

			// Merge field and answer if assignment is delivered
			for i := 0; i < len(form.Form.Fields); i++ {
				com.Items = append(com.Items, Combined{
					Answer: assignmentAnswers[i],
					Field:  form.Form.Fields[i],
				})
			}
		}
	*/

	// Get review form for the assignment
	review, err := services.Review.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get single review", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header content-type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "assignment/submission"
	// View variables
	v.Vars["Assignment"] = assignment
	v.Vars["User"] = user
	v.Vars["Course"] = course
	v.Vars["Delivered"] = len(assignmentAnswers) > 0
	v.Vars["IsTeacher"] = currentUser.Teacher
	v.Vars["Fields"] = form.Form.Fields
	v.Vars["Answers"] = assignmentAnswers
	if review.FormID != 0 {
		v.Vars["Review"] = review
	}

	// Render view
	v.Render(w)
}

// AssignmentUserSubmissionPOST handles POST-request @ /assignment/{id:[0-9]+}/submission/{userid:[0-9]+}
func AssignmentUserSubmissionPOST(w http.ResponseWriter, r *http.Request) {
	currentUser := session.GetUserFromSession(r)
	if !currentUser.Authenticated {
		log.Printf("Error: Could not get user (assignment.go)")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	p := bluemonday.UGCPolicy()

	assignmentID, err := strconv.Atoi(p.Sanitize(vars["id"]))
	if err != nil {
		log.Println("id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	targetID, err := strconv.Atoi(p.Sanitize(vars["userid"]))
	if err != nil {
		log.Println("userid", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviewID, err := strconv.Atoi(p.Sanitize(r.FormValue("review_id")))
	if err != nil {
		log.Println("review_id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	hasBeenReviewed, err := services.ReviewAnswer.HasBeenReviewed(targetID, currentUser.ID, assignmentID)
	if err != nil {
		log.Println("has been reviewed", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if hasBeenReviewed {
		IndexGET(w, r)
		return
	}

	// Parse form from the request
	err = r.ParseForm()
	if err != nil {
		log.Println("parse form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	form, err := services.Review.FetchFromAssignment(assignmentID)
	if err != nil {
		log.Println("get review form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviewAnswer := make([]model.ReviewAnswer, 0)

	for _, field := range form.Form.Fields {
		temp := model.ReviewAnswer{
			UserReviewer: currentUser.ID,
			UserTarget:   targetID,
			AssignmentID: assignmentID,
			ReviewID:     reviewID,
			Type:         field.Type,
			Name:         field.Name,
			Label:        field.Label,
			Description:  field.Description,
			Answer:       p.Sanitize(r.FormValue(field.Name)),
			HasComment:   field.HasComment,
			Choices:      field.Choices,
			Weight:       field.Weight,
		}

		if field.HasComment {
			comment := p.Sanitize(r.FormValue(field.Name + "_comment"))
			temp.Comment = sql.NullString{
				String: comment,
				Valid:  comment != "",
			}
		}

		reviewAnswer = append(reviewAnswer, temp)
	}

	/*
	fullReview := model.FullReview{
		Reviewer:     currentUser.ID,
		Target:       targetID,
		ReviewID:     reviewID,
		AssignmentID: assignmentID,
		Answers:      make([]model.ReviewAnswer, 0),
	}

	for _, field := range form.Form.Fields {
		answer := model.ReviewAnswer{
			Type:   field.Type,
			Name:   field.Name,
			Label:  field.Label,
			Answer: p.Sanitize(r.FormValue(field.Name)),
		}

		fullReview.Answers = append(fullReview.Answers, answer)
	}
	*/

	for _, item := range reviewAnswer {
		_, err = services.ReviewAnswer.Insert(item)
		if err != nil {
			log.Println("services, review answer, insert", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// TODO (Svein): Want to send back to /assignment/{id}. HOW TO?
	IndexGET(w, r)
}
