package model

import (
	"database/sql"
	"time"
)

// Activity enum for keeping track of user activity
type Activity string

// Enum for logs
const (
	/////// System logs ///////
	NewUser             Activity = "NEW-USER"                      // A new user is created
	ChangeEmail         Activity = "CHANGE-EMAIL"                  // User changed email
	ChangePassword      Activity = "CHANGE-PASSWORD"               // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	ChangePasswordEmail Activity = "CHANGE-PASSWORD-THROUGH-EMAIL" // User changed password through email

	/////// Course logs ///////
	// Submission
	DeliveredSubmission Activity = "SUBMISSION-DELIVERED" // User delivered assignment submission
	UpdateSubmission    Activity = "SUBMISSION-UPDATED"   // User re-delivered assignment submission
	// Review
	FinishedOnePeerReview Activity = "FINISHED-ONE-PEER-REVIEW" // User is done with one peer review (that this user did)
	UpdateOnePeerReview   Activity = "UPDATE-ONE-PEER-REVIEW"   // User changed one peer review
	// Course
	JoinedCourse Activity = "JOINED-COURSE" // User joined course
	LeftCourse   Activity = "LEFT-COURSE"   // USer left course

	/////// Admin logs ///////
	// Assignment
	AdminCreatAssignment  Activity = "ADMIN-ASSIGNMENT-CREATE" // Admin assignment is created
	AdminDeleteAssignment Activity = "ADMIN-ASSIGNMENT-DELETE" // Admin assignment is deleted
	AdminUpdateAssignment Activity = "ADMIN-ASSIGNMENT-UPDATE" // Admin assignment is updated
	// Forms
	AdminCreateSubmissionForm Activity = "ADMIN-SUBMISSION-FORM-CREATE" // Admin submission form is created
	AdminUpdateSubmissionForm Activity = "ADMIN-SUBMISSION-FORM-UPDATE" // Admin submission form is updated
	AdminDeleteSubmissionForm Activity = "ADMIN-SUBMISSION-FORM-DELETE" // Admin submission form is deleted
	AdminCreateReviewForm     Activity = "ADMIN-REVIEW-FORM-CREATE"     // Admin review form is created
	AdminUpdateReviewForm     Activity = "ADMIN-REVIEW-FORM-UPDATE"     // Admin review form is updated
	AdminDeleteReviewForm     Activity = "ADMIN-REVIEW-FORM-DELETE"     // Admin review form is deleted
	// Course
	AdminCreatedCourse Activity = "ADMIN-COURSE-CREATE" // Admin course is created
	AdminUpdateCourse  Activity = "ADMIN-COURSE-UPDATE" // Admin course is updated
	AdminDeleteCourse  Activity = "ADMIN-COURSE-DELETE" // Admin course is deleted
	// FAQ
	AdminCreateFAQ Activity = "ADMIN-CREATE-FAQ" // Admin FAQ is created
	AdminUpdateFAQ Activity = "ADMIN-UPDATE-FAQ" // Admin FAQ is updated
	AdminDeleteFAQ Activity = "ADMIN-DELETE-FAQ" // Admin FAQ is deleted
	// Manage students
	AdminEmailCourseStudents     Activity = "ADMIN-EMAIL-STUDENTS-IN-COURSE"   // Admin emailed all students in course through the system
	AdminRemoveUserFromCourse    Activity = "ADMIN-REMOVE-USER-FROM-COURSE"    // Admin removed one user from course
	AdminChangeStudentPassword   Activity = "ADMIN-CHANGE-STUDENT-PASSWORD"    // Admin changed one users password
	AdminCreateSubmissionForUser Activity = "ADMIN-CREATE-SUBMISSION-FOR-USER" // Admin created submission for user
	AdminUpdateSubmissionForUser Activity = "ADMIN-UPDATE-SUBMISSION-FOR-USER" // Admin updated submission for user
	AdminDeleteSubmissionForUser Activity = "ADMIN-DELETE-SUBMISSION-FOR-USER" // Admin deleted submission for user
)

// Logs struct for keeping logs data
type Logs struct {
	ID             int            `json:"id"`               // [NOT NULL][all]
	UserID         int            `json:"user_id"`          // [NOT NULL][all]
	Timestamp      time.Time      `json:"timestamp"`        // [NOT NULL][all]
	Activity       Activity       `json:"activity"`         // [NOT NULL][all]
	AssignmentID   sql.NullInt64  `json:"assignment_id"`    // [NULLABLE]
	CourseID       sql.NullInt64  `json:"course_id"`        // [NULLABLE]
	SubmissionID   sql.NullInt64  `json:"submission_id"`    // [NULLABLE]
	ReviewID       sql.NullInt64  `json:"review_id"`        // [NULLABLE]
	OldValue       sql.NullString `json:"old_value"`        // [NULLABLE]
	NewValue       sql.NullString `json:"new_value"`        // [NULLABLE]
	AffectedUserID sql.NullInt64  `json:"affected_user_id"` // [NULLABLE]
}
