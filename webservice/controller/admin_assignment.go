package controller

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/scheduler"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// AdminAssignmentGET handles GET-request at /admin/assignment
func AdminAssignmentGET(w http.ResponseWriter, r *http.Request) {

	// repo's
	courseRepo := &model.CourseRepository{}
	assignmentRepo := model.AssignmentRepository{}

	//get courses to user/teacher
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//need custom struct to get the coursecode
	type ActiveAssignment struct {
		Assignment model.Assignment
		CourseCode string
	}

	var activeAssignments []ActiveAssignment

	for _, course := range courses { //iterate all courses
		assignments, err := assignmentRepo.GetAllFromCourse(course.ID) //get assignments from course
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		for _, assignment := range assignments { //go through all it's assignments again
			activeAssignments = append(activeAssignments, ActiveAssignment{Assignment: assignment, CourseCode: course.Code})
		}

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/index"

	v.Vars["Assignments"] = activeAssignments

	v.Render(w)
}

// AdminAssignmentCreateGET handles GET-request from /admin/assigment/create
func AdminAssignmentCreateGET(w http.ResponseWriter, r *http.Request) {
	var subRepo = model.SubmissionRepository{}
	submissions, err := subRepo.GetAll()

	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviewRepo := model.ReviewRepository{}
	reviews, err := reviewRepo.GetAll()

	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//course repo
	courseRepo := model.CourseRepository{}
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)

	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/create"

	v.Vars["Courses"] = courses

	v.Vars["Submissions"] = submissions
	v.Vars["Reviews"] = reviews

	v.Render(w)
}

// AdminAssignmentCreatePOST handles POST-request from /admin/assigment/create
func AdminAssignmentCreatePOST(w http.ResponseWriter, r *http.Request) {

	// Declare empty slice of strings
	var errorMessages []string

	// Get form name from request
	assignmentName := r.FormValue("name")
	// Get form description from request
	assignmentDescription := r.FormValue("description")

	// Check if name is empty
	if assignmentName == "" {
		errorMessages = append(errorMessages, "Error: Assignment Name cannot be blank.")
	}

	// Get the time.Time object from the publish string
	publish, err := util.DatetimeLocalToRFC3339(r.FormValue("publish"))
	if err != nil {
		errorMessages = append(errorMessages, "Error: Something wrong with the publish datetime.")
	}

	// Get the time.Time object from the deadline string
	deadline, err := util.DatetimeLocalToRFC3339(r.FormValue("deadline"))
	if err != nil {
		errorMessages = append(errorMessages, "Error: Something wrong with the deadline datetime.")
	}

	// Check if publish datetime is after the deadline
	if publish.After(deadline) {
		errorMessages = append(errorMessages, "Error: Deadline cannot be before Publish.")
	}

	// Check if there are any error messages
	if len(errorMessages) != 0 {
		// TODO (Svein): Keep data from the previous submit
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		//course repo
		courseRepo := &model.CourseRepository{}

		courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		v := view.New(r)
		v.Name = "admin/assignment/create"

		v.Vars["Errors"] = errorMessages
		v.Vars["AssignmentName"] = assignmentName
		v.Vars["AssignmentDescription"] = assignmentDescription
		v.Vars["Courses"] = courses

		v.Render(w)
		return
	}

	assignmentRepository := model.AssignmentRepository{}
	// Get form values
	var val int

	// String converted into integer
	courseID, err := strconv.Atoi(r.FormValue("course_id"))
	if err != nil {
		log.Print("course_id")
		log.Println(err)
		return
	}

	if r.FormValue("submission_id") != "" {
		val, err = strconv.Atoi(r.FormValue("submission_id"))
		if err != nil {
			log.Println("submission_id")
			log.Println(err)
			return
		}
	}
	submissionID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	val = 0

	if r.FormValue("review_id") != "" {
		val, err = strconv.Atoi(r.FormValue("review_id"))
		if err != nil {
			log.Println("review_id")
			log.Println(err)
			return
		}
	}
	reviewID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	if r.FormValue("reviewers") != "" {
		val, err = strconv.Atoi(r.FormValue("reviewers"))
		if err != nil {
			log.Println("reviewers")
			log.Println(err)
			return
		}
	}
	reviewers := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	// Put all data into an Assignment-struct
	assignment := model.Assignment{
		Name:         assignmentName,
		Description:  assignmentDescription,
		Publish:      publish,
		Deadline:     deadline,
		CourseID:     courseID,
		SubmissionID: submissionID,
		ReviewID:     reviewID,
		Reviewers:    reviewers,
	}

	// Insert data to database
	assID, err := assignmentRepository.Insert(assignment)
	if err != nil {
		log.Println(err)
		return
	}

	// if submission ID AND Reviewers is set and valid, we can schedule the peer_review service to execute
	if assID != 0 && assignment.SubmissionID.Valid && assignment.Reviewers.Valid && assignment.Deadline.After(time.Now()) {

		sched := scheduler.Scheduler{}

		err := sched.SchedulePeerReview(int(assignment.SubmissionID.Int64),
			assID, //assignment ID
			int(assignment.Reviewers.Int64),
			assignment.Deadline)
		if err != nil {
			log.Println(err)
			return
		}

	}

	http.Redirect(w, r, "/admin/assignment", http.StatusFound)
}

// AdminSingleAssignmentGET handles GET-request at admin/assignment/{id}
func AdminSingleAssignmentGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	assignmentRepo := &model.AssignmentRepository{}

	assignment, err := assignmentRepo.GetSingle(int(id))
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	descriptionMD := []byte(assignment.Description)
	description := github_flavored_markdown.Markdown(descriptionMD)

	v := view.New(r)
	v.Name = "admin/assignment/single"

	v.Vars["Assignment"] = assignment
	v.Vars["Description"] = template.HTML(description)

	v.Render(w)
}

// AdminUpdateAssignmentGET handles GET-request at /admin/assignment/update/{id}
func AdminUpdateAssignmentGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	assignmentRepo := &model.AssignmentRepository{}
	submissionRepo := &model.SubmissionRepository{}
	reviewRepo := &model.ReviewRepository{}
	courseRepo := &model.CourseRepository{}

	submissions, err := submissionRepo.GetAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignment, err := assignmentRepo.GetSingle(int(id))
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get number of Students that has delivered submission with specific submission form
	submissionCount, err := submissionRepo.GetSubmissionsCountFromAssignment(assignment.ID, assignment.SubmissionID.Int64)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//get courses to user
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviews, err := reviewRepo.GetAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	v := view.New(r)
	v.Name = "admin/assignment/update"

	v.Vars["Assignment"] = assignment
	v.Vars["SubmissionCount"] = submissionCount
	v.Vars["Publish"] = util.GoToHTMLDatetimeLocal(assignment.Publish)
	v.Vars["Deadline"] = util.GoToHTMLDatetimeLocal(assignment.Deadline)
	v.Vars["Courses"] = courses
	v.Vars["Submissions"] = submissions
	v.Vars["Reviews"] = reviews

	v.Render(w)
}

// AdminUpdateAssignmentPOST handles POST-request at /admin/assignment/update
func AdminUpdateAssignmentPOST(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get the time.Time object from the publish string
	publish, err := util.DatetimeLocalToRFC3339(r.FormValue("publish"))
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get the time.Time object from the deadline string
	deadline, err := util.DatetimeLocalToRFC3339(r.FormValue("deadline"))
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if publish datetime is after the deadline
	if publish.After(deadline) {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form values
	courseIDString := r.FormValue("course_id")
	var val int

	// String converted into integer
	courseID, err := strconv.Atoi(courseIDString)
	if err != nil {
		log.Printf("course_id: %v", err)
		return
	}

	if r.FormValue("submission_id") != "" {
		val, err = strconv.Atoi(r.FormValue("submission_id"))
		if err != nil {
			log.Println("submission_id")
			log.Println(err)
			return
		}
	}

	// TODO brede : I think I need the value to be 0 sometimes
	submissionID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	val = 0

	if r.FormValue("review_id") != "" {
		val, err = strconv.Atoi(r.FormValue("review_id"))
		if err != nil {
			log.Println("review_id")
			log.Println(err)
			return
		}
	}
	reviewID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	if r.FormValue("reviewers") != "" {
		val, err = strconv.Atoi(r.FormValue("reviewers"))
		if err != nil {
			log.Println("reviewers")
			log.Println(err)
			return
		}
	}
	reviewers := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	assignmentRepo := model.AssignmentRepository{}

	assignment := model.Assignment{
		ID:           id,
		Name:         r.FormValue("name"),
		Description:  r.FormValue("description"),
		Publish:      publish,
		Deadline:     deadline,
		CourseID:     courseID,
		SubmissionID: submissionID,
		ReviewID:     reviewID,
		Reviewers:    reviewers,
	}

	err = assignmentRepo.Update(id, assignment)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// if submission ID AND Reviewers is set and valid, we can schedule the peer_review service to execute
	if assignment.ID != 0 && assignment.SubmissionID.Valid && assignment.Reviewers.Valid && assignment.Deadline.After(time.Now()) {

		sched := scheduler.Scheduler{}

		if sched.SchedulerExists(int(assignment.SubmissionID.Int64), assignment.ID) {
			err := sched.UpdateSchedule(int(assignment.SubmissionID.Int64),
				assignment.ID, //assignment ID
				assignment.Deadline)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			err := sched.SchedulePeerReview(int(assignment.SubmissionID.Int64),
				assignment.ID, //assignment ID
				int(assignment.Reviewers.Int64),
				assignment.Deadline)
			if err != nil {
				log.Println(err)
				return
			}
		}

	}

	http.Redirect(w, r, "/admin/assignment", http.StatusFound)
}

// AdminAssignmentSubmissionsGET servers list of all users in course to admin
func AdminAssignmentSubmissionsGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignmentRepo := &model.AssignmentRepository{}

	assignment, err := assignmentRepo.GetSingle(int(assignmentID))
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	courseRepo := &model.CourseRepository{}

	course, err := courseRepo.GetSingle(assignment.CourseID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	submissionRepo := &model.SubmissionRepository{}

	submissionCount, err := submissionRepo.GetSubmissionsCountFromAssignment(assignment.ID, assignment.SubmissionID.Int64)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO brede : sort by user delivered and not + show if delivered or not in table
	students := model.GetUsersToCourse(assignment.CourseID)
	if len(students) < 0 {
		log.Println("Error: could not get students from course! (admin_assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	type UserAndSubmit struct {
		User          model.User
		SubmittedTime time.Time
		Submitted     bool
	}

	var users []UserAndSubmit

	for _, student := range students {
		submitTime, submitted, err := model.GetSubmittedTime(student.ID, assignmentID)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		if submitted {
			var data = UserAndSubmit{
				User:          student,
				SubmittedTime: submitTime,
				Submitted:     true,
			}

			users = append(users, data)
		} else {
			var data = UserAndSubmit{
				User:      student,
				Submitted: false,
			}

			users = append(users, data)
		}
	}

	//Sort slice by submitted time
	sort.Slice(users, func(i, j int) bool {
		return users[i].SubmittedTime.After(users[j].SubmittedTime)
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/submissions"

	v.Vars["SubmissionCount"] = submissionCount
	v.Vars["Assignment"] = assignment
	v.Vars["Students"] = users
	v.Vars["Course"] = course

	v.Render(w)
}

/*
TODO brede : use this with iframe after alpha
// AdminAssignmentSubmissionGET servers one user submission in course to admin
func AdminAssignmentSubmissionGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("userid"))
	if err != nil {
		log.Printf("userid: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	fmt.Println(userID) // TODO brede : remove this

	assignmentRepo := &model.AssignmentRepository{}

	assignment, err := assignmentRepo.GetSingle(int(assignmentID))
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO brede : use same page as peer rews aka. out of admin/
	v := view.New(r)
	v.Name = "admin/assignment/singlesubmission"

	v.Vars["Assignment"] = assignment

	v.Render(w)

}
*/
