package model

import (
	"database/sql"
	"time"
)

// Activity enum for keeping track of user activity
type Activity int

// Enum for logs
const (
	/////// System logs ///////
	NewUser             Activity = 0 // A new user is created
	ChangeEmail         Activity = 1 // User changed email
	ChangePassword      Activity = 2 // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	ChangePasswordEmail Activity = 3 // User changed password through email

	/////// Course logs ///////
	// Submission
	CreateSubmission Activity = 4 // User submission is created
	UpdateSubmission Activity = 5 // User submission is updated
	DeleteSubmission Activity = 6 // User submission is deleted
	// Review
	FinishedOnePeerReview Activity = 7 // User is done with one peer review (that this user did)
	UpdateOnePeerReview   Activity = 8 // User changed one peer review
	// Course
	JoinedCourse Activity = 9  // User joined course
	LeftCourse   Activity = 10 // USer left course
	// Group
	CreateGroup     Activity = 11 // User created group
	EditGroupName   Activity = 12 // User edited group name
	DeleteGroup     Activity = 13 // User deleted group
	JoinGroup       Activity = 14 // User joined group
	LeftGroup       Activity = 15 // User left group
	KickedFromGroup Activity = 16 // User kicked from group

	/////// Admin logs ///////
	// Assignment
	AdminCreateAssignment Activity = 100 // Admin assignment is created
	AdminDeleteAssignment Activity = 101 // Admin assignment is deleted
	AdminUpdateAssignment Activity = 102 // Admin assignment is updated
	// Forms
	AdminCreateSubmissionForm Activity = 103 // Admin submission form is created
	AdminUpdateSubmissionForm Activity = 104 // Admin submission form is updated
	AdminDeleteSubmissionForm Activity = 105 // Admin submission form is deleted
	AdminCreateReviewForm     Activity = 106 // Admin review form is created
	AdminUpdateReviewForm     Activity = 107 // Admin review form is updated
	AdminDeleteReviewForm     Activity = 108 // Admin review form is deleted
	// Course
	AdminCreatedCourse Activity = 109 // Admin course is created
	AdminUpdateCourse  Activity = 110 // Admin course is updated
	AdminDeleteCourse  Activity = 111 // Admin course is deleted
	// FAQ
	AdminCreateFAQ Activity = 112 // Admin FAQ is created
	AdminUpdateFAQ Activity = 113 // Admin FAQ is updated
	// Manage students
	AdminEmailCourseStudents     Activity = 114 // Admin emailed all students in course through the system
	AdminRemoveUserFromCourse    Activity = 115 // Admin removed one user from course
	AdminChangeStudentPassword   Activity = 116 // Admin changed one users password
	AdminCreateSubmissionForUser Activity = 117 // Admin created submission for user
	AdminUpdateSubmissionForUser Activity = 118 // Admin updated submission for user
	AdminDeleteSubmissionForUser Activity = 119 // Admin deleted submission for user
	AdminAddUserToGroup          Activity = 120 // Admin added user to group
	AdminRemoveUserFromGroup     Activity = 121 // Admin removed user from group
	AdminEditGroupName           Activity = 122 // Admin edited group name
	AdminDeleteGroup             Activity = 123 // Admin deleted group
	AdminCreateGroup             Activity = 124 // Admin created group
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
	GroupID        sql.NullInt64  `json:"group_id"`         // [NULLABLE]
	OldValue       sql.NullString `json:"old_value"`        // [NULLABLE]
	NewValue       sql.NullString `json:"new_value"`        // [NULLABLE]
	AffectedUserID sql.NullInt64  `json:"affected_user_id"` // [NULLABLE]
}
