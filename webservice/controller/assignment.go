package controller

import (
	"database/sql"
	"fmt"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	// Filter out the reviews that the current user already has done
	reviewUsers, err := services.Review.FetchReviewUsers(currentUser.ID, assignment.ID)
	if err != nil {
		log.Println("services, review, fetch review users", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	/*
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
	*/

	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("services, course, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviews, err := services.ReviewAnswer.FetchForUser(currentUser.ID, assignmentID)
	if err != nil {
		log.Println("review answer service, fetch for target", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO time-norwegian
	var isDeadlineOver = assignment.Deadline.Before(util.GetTimeInCorrectTimeZone())

	var isReviewDeadlineOver = assignment.ReviewDeadline.Before(util.GetTimeInCorrectTimeZone())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Create view
	v := view.New(r)

	// Check if assignment has group delivery
	if assignment.GroupDelivery {
		v.Name = "assignment/index_group"

		// Check if current user has group
		hasGroup, err := services.GroupService.UserInAnyGroup(currentUser.ID, assignment.ID)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Fetch all groups
		groups, err := services.GroupService.FetchAll(assignment.ID)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		v.Vars["HasGroup"] = hasGroup
		v.Vars["Groups"] = groups

		// Check if current user has a group
		if hasGroup {
			// Fetch group for current user
			group, err := services.GroupService.FetchGroupForUser(currentUser.ID, assignment.ID)
			if err != nil {
				log.Println(err.Error())
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
			// Fetch group members for group
			groupMembers, err := services.GroupService.FetchUsersFromGroup(group.ID)
			if err != nil {
				log.Println(err.Error())
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
			// Put members into group
			for _, user := range groupMembers {
				group.Users = append(group.Users, *user)
			}

			v.Vars["Group"] = group
		}
	} else { // Not group delivery
		v.Name = "assignment/index"
	}

	v.Vars["Message"] = session.GetFlash(w, r)
	v.Vars["Assignment"] = assignment
	v.Vars["Delivered"] = delivered
	v.Vars["HasReview"] = hasReview
	v.Vars["IsDeadlineOver"] = isDeadlineOver
	v.Vars["IsReviewDeadlineOver"] = isReviewDeadlineOver
	v.Vars["CourseID"] = course.ID
	v.Vars["Reviews"] = reviewUsers
	v.Vars["MyReviews"] = reviews
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

	if !assignment.SubmissionID.Valid {
		log.Println("assignment submission_id is not valid, redirecting user")
		http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignmentID), http.StatusTemporaryRedirect)
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
			answers[index].Required = field.Required
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
				Required:    item.Required,
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sess, err := session.Instance(r)
	if err != nil {
		log.Println("session, instance", err)
	}

	var successMessage string

	successFlash := sess.Flashes("success")
	if len(successFlash) > 0 {
		successMessage = successFlash[0].(string)
	}

	// Create view
	v := view.New(r)
	v.Name = "assignment/upload"

	v.Vars["Course"] = course
	v.Vars["Assignment"] = assignment
	v.Vars["Fields"] = submissionForm.Form.Fields
	v.Vars["Delivered"] = delivered
	v.Vars["Answers"] = answers
	v.Vars["SuccessMessage"] = successMessage

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
	services := service.NewServices(db.GetDB())

	// Get assignment and log possible error
	assignment, err := services.Assignment.Fetch(assignmentID)
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

	// TODO time-norwegian
	loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		log.Println(err.Error())
	}

	// TODO fix hack
	deadline := assignment.Deadline.In(loc).Add(-time.Hour)

	// Check if the deadline is reached TODO fix this quick fix time-norwegian

	var isDeadlineOver = deadline.Before(util.GetTimeInCorrectTimeZone())
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
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("submission service, fetch from assignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if user has uploaded already or not
	delivered, err := services.SubmissionAnswer.HasUserSubmitted(assignmentID, currentUser.ID)
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
		submissionAnswers, err = services.SubmissionAnswer.FetchUserAnswers(currentUser.ID, assignment.ID)
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

	// Insert or update answers
	if !delivered {

		// Insert new answer
		err = services.SubmissionAnswer.Insert(submissionAnswers)
		if err != nil {
			log.Println("submission answer service, upload", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Log inserting of new submission
		err = services.Logs.InsertSubmission(currentUser.ID, assignment.ID, int(assignment.SubmissionID.Int64))
		if err != nil {
			log.Println("log, insert new submission", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	} else {

		// Update answer
		err = services.SubmissionAnswer.Update(submissionAnswers)
		if err != nil {
			log.Println("submission answer service, update", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Log inserting of updated submission
		err = services.Logs.InsertUpdateSubmission(currentUser.ID, assignment.ID, int(assignment.SubmissionID.Int64))
		if err != nil {
			log.Println("log, update submission", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	sess, err := session.Instance(r)
	if err != nil {
		log.Println("session, instance", err)
	}

	sess.AddFlash("Submission submitted!", "success")

	// Serve front-end again
	AssignmentUploadGET(w, r)
}

// AssignmentUserSubmissionGET serves one user submission to admin and the peer reviews
func AssignmentUserSubmissionGET(w http.ResponseWriter, r *http.Request) {
	// Services
	services := service.NewServices(db.GetDB())

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
	user, err := services.User.Fetch(userID)
	if err != nil {
		log.Printf("Error: Could not get user (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Current user
	currentUser := session.GetUserFromSession(r)

	// Get relevant assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("get single assignment", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Review deadline is zero, send user to front-page
	if assignment.ReviewDeadline.IsZero() {
		log.Println("reviewDeadline isZero")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Check review deadline
	if assignment.ReviewDeadline.Before(util.GetTimeInCorrectTimeZone()) {
		log.Println("DEBUG:", assignment.ReviewDeadline.UTC(), "after", util.GetTimeInCorrectTimeZone())
		ErrorHandler(w, r, http.StatusTeapot)
		return
	}

	// Give error if user isn't teacher or reviewer for this user
	isUserTheReviewer, err := services.Review.IsUserTheReviewer(currentUser.ID, userID, assignment.ID)
	if err != nil {
		log.Println("services, review, is user the reviewer", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	isUserTheOwner := false
	if userID == currentUser.ID {
		isUserTheOwner = true
	}

	//if !currentUser.Teacher && !model.UserIsReviewer(currentUser.ID, assignment.ID, assignment.SubmissionID.Int64, userID) {
	if !currentUser.Teacher && !isUserTheReviewer && !isUserTheOwner {
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

	// Get review form for the assignment
	reviewForm, err := services.Review.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get single review", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	review, err := services.ReviewAnswer.FetchForReviewerAndTarget(currentUser.ID, userID, assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if len(review) == 0 {
		for _, field := range reviewForm.Form.Fields {
			ra := model.ReviewAnswer{}

			ra.Name = field.Name
			ra.Type = field.Type
			ra.Label = field.Label
			ra.Description = field.Description
			ra.Choices = field.Choices
			ra.Required = field.Required
			ra.HasComment = field.HasComment

			review = append(review, &ra)
		}
	}

	reviewMessages, err := services.ReviewMessage.FetchAllForAssignmentUser(assignmentID, userID)
	if err != nil {
		log.Println("AssignmnetUserSubmissionGET, fetch review messages, ", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header content-type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
	v.Vars["IsOwner"] = isUserTheOwner
	v.Vars["ReviewMessages"] = reviewMessages

	v.Vars["Review"] = review

	// Render view
	v.Render(w)
}

// AssignmentUserSubmissionPOST handles POST-request @ /assignment/{id:[0-9]+}/submission/{userid:[0-9]+}
func AssignmentUserSubmissionPOST(w http.ResponseWriter, r *http.Request) {
	currentUser := session.GetUserFromSession(r)

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

	// Services
	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	hasBeenReviewed, err := services.ReviewAnswer.HasBeenReviewed(targetID, currentUser.ID, assignmentID)
	if err != nil {
		log.Println("has been reviewed", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
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

	if hasBeenReviewed {

		// Update answers/submission
		for _, field := range form.Form.Fields {
			answer := r.FormValue(field.Name)
			comment := r.FormValue(field.Name + "_comment")

			err = services.ReviewAnswer.Update(targetID, currentUser.ID, assignmentID, model.ReviewAnswer{
				Answer: answer,
				Comment: sql.NullString{
					String: comment,
					Valid:  comment != "",
				},
				Name: field.Name,
			})
			if err != nil {
				log.Println(err.Error())
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}

		// Log updated review
		err = services.Logs.InsertUpdateOnePeerReview(currentUser.ID, assignmentID, targetID)
		if err != nil {
			log.Println("log, update review", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignment.ID), http.StatusFound)
		return
	}

	reviewAnswer := make([]model.ReviewAnswer, 0)

	for _, field := range form.Form.Fields {
		temp := model.ReviewAnswer{
			UserReviewer: currentUser.ID,
			UserTarget:   targetID,
			AssignmentID: assignmentID,
			ReviewID:     int(assignment.ReviewID.Int64),
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

	// Insert answers
	for _, item := range reviewAnswer {
		_, err = services.ReviewAnswer.Insert(item)
		if err != nil {
			log.Println("services, review answer, insert", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	// Log new review
	err = services.Logs.InsertFinishedOnePeerReview(currentUser.ID, assignmentID, targetID)
	if err != nil {
		log.Println("log, update review", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignment.ID), http.StatusFound)
}

// AssignmentWithdrawGET handles GET-requests for withdrawing submissions
func AssignmentWithdrawGET(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, atoi, id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Services
	services := service.NewServices(db.GetDB())

	// Delete user submission
	err = services.SubmissionAnswer.Delete(assignmentID, currentUser.ID)
	if err != nil {
		log.Println("services, submission answer, delete", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch submission id for logging purposes
	submission, err := services.Submission.FetchFromAssignment(assignmentID)
	if err != nil {
		log.Println("services, submission, fetchfromassignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log deletion of submission
	err = services.Logs.InsertDeleteSubmission(currentUser.ID, assignmentID, submission.ID)
	if err != nil {
		log.Println("log, delete submission", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	sess, err := session.Instance(r)
	if err != nil {
		log.Println("session, instace", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	sess.AddFlash("Submission withdrawn!", "success")

	IndexGET(w, r)
}

// AssignmentReviewRequestPOST requests a new review to be assigned
func AssignmentReviewRequestPOST(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("AssignmentReviewRequestPOST, strconv, atoi, id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Services
	services := service.NewServices(db.GetDB())

	//get assignment info
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("AssignmentReviewRequestPOST, services.Assignment.Fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//check that user is asking before the deadline
	if assignment.ReviewDeadline.Before(util.GetTimeInCorrectTimeZone()) {
		log.Println("AssignmentReviewRequestPOST, assignment.ReviewDeadline.Before", err)
		ErrorHandler(w, r, http.StatusTeapot)
		return
	}

	usersDelivered, err := services.SubmissionAnswer.FetchUsersDeliveredFromAssignment(assignmentID)
	if err != nil {
		log.Println("AssignmentReviewRequestPOST, services.Submission.FetchFromAssignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//get all the users that the user is currently "reviewing"
	alreadyTargets, err := services.PeerReview.FetchReviewTargetsToUser(currentUser.ID, assignmentID)
	if err != nil {
		log.Println("AssignmentReviewRequestPOST, services.Submission.FetchReviewTargetsToUser", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//find the lowest amount of reviews
	lowestNrReviews := 99999
	usersAndReviews := make(map[int]int)

	//set the review count on already reviewed users very high
	for _, target := range alreadyTargets {
		usersAndReviews[target.TargetID] = 99999
	}

	for _, user := range usersDelivered {
		if user != currentUser.ID && usersAndReviews[user] != 99999 { //don't include self or already reviewed targets
			reviewsDone, err := services.ReviewAnswer.FetchForTarget(user, assignmentID)
			if err != nil {
				log.Println("AssignmentReviewRequestPOST, services.ReviewAnswer.FetchForTarget", err)
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}

			usersAndReviews[user] = len(reviewsDone)

			if len(reviewsDone) < lowestNrReviews {
				lowestNrReviews = len(reviewsDone)
			}
		}
	}

	//if you have reviewed everyone
	if lowestNrReviews > 99998 {
		_ = session.SetFlash("Review Limit Reached: You have reviewed every possible submission.", w, r)
		AssignmentSingleGET(w, r)

		return
	}

	//filter the submissions with lowest reviewcount
	submissionsFiltered := make([]int, 0)

	for userID, reviews := range usersAndReviews {
		if reviews == lowestNrReviews {
			submissionsFiltered = append(submissionsFiltered, userID)
		}
	}

	//shuffle slice
	submissionsFiltered = util.ShuffleIntSlice(submissionsFiltered)

	//save the 0 index as a new review pair
	inserted, err := services.PeerReview.Insert(assignmentID, currentUser.ID, submissionsFiltered[0])
	if err != nil {
		log.Println("AssignmentReviewRequestPOST, services.PeerReview.Insert", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//redirect
	if inserted {
		http.Redirect(w, r, fmt.Sprintf("/assignment/%v/submission/%v", assignmentID, submissionsFiltered[0]), http.StatusFound)
	} else {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// AssignmentUserSubmissionMessagePOST function saves message and creates notification
func AssignmentUserSubmissionMessagePOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()

	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, atoi, assignmentID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Println("strconv, atoi, userID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//fetch URL get parameters
	getKeys := r.URL.Query()

	replyID := 0
	replyStringID := getKeys.Get("reply")

	if replyStringID != ""{
		replyID, err = strconv.Atoi(replyStringID)
		if err != nil {
			log.Println("strconv, atoi, replyStringID", err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	notifyAll := false
	notifyAllString := getKeys.Get("notifyAll")

	if notifyAllString != ""{
		notifyAll = true
	}

	// Services
	services := service.NewServices(db.GetDB())

	// double check that messages are enabled in assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("assignmentUserSubmissionMessagePost, serivces, assignmetn, fetc, ", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if !assignment.ReviewEnabled {
		log.Println("reviewMessages is not enabled")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Get message from form
	message := p.Sanitize(r.FormValue("message_field"))

	//reviewMessage object
	reviewMessage := model.ReviewMessage{}
	//fill with data
	reviewMessage.AssignmentID = assignmentID
	reviewMessage.UserTarget = userID
	reviewMessage.UserReviewer = currentUser.ID
	reviewMessage.Message = message

	err = services.ReviewMessage.Insert(reviewMessage)
	if err != nil {
		log.Println("reviewMessage, insert, ", err)
		//todo use log service
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//dont notify yourself
	if userID != currentUser.ID && replyID != 0 {
		//creating notification for owner of the assignment
		notification := model.Notification{}
		notification.UserID = replyID
		notification.Message = "There has been posted a reply on your comment."
		notification.URL = fmt.Sprintf("/assignment/%v/submission/%v", assignmentID, userID)

		_, err = services.Notification.Insert(notification)
		if err != nil {
			log.Println("reviewMessage, notification, insert ", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
		}
	}

	if notifyAll{
		//todo fetch all reviewers on submission and send them a notification

	}

	//redirect back
	http.Redirect(w, r, fmt.Sprintf("/assignment/%v/submission/%v", assignmentID, userID), http.StatusFound)
}
