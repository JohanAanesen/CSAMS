package model

import (
	"database/sql"
	"time"
)

// Activity enum for keeping track of log Activity
type Activity string

// Enum for logs
const (
	NewUser               Activity = "NEW-USER"                               // A new user is created
	ChangeEmail           Activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword        Activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment   Activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	UpdateAssignment      Activity = "ASSIGNMENT-UPDATED"                     // User re-delivered assignment
	DeleteAssignment      Activity = "ASSIGNMENT-DELETE"                      // user deleted assignment
	FinishedOnePeerReview Activity = "FINISHED-ONE-PEER-REVIEW"               // User is done with one peer review (that this user did)
	PeerReviewDone        Activity = "PEER-REVIEW-IS-DONE-FOR-ONE-SUBMISSION" // User submission is finished peer-reviewed by other students
	JoinedCourse          Activity = "JOINED-COURSE"                          // User joined course
	CreatedCourse         Activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment       Activity = "ASSIGNMENT-CREATED"                     // Assignment is created
	ChangeAdminFAQ        Activity = "CHANGE-ADMIN-FAQ"                       // The admins faq is updated
	// TODO Brede : add more activities later :)
)

// Logs struct for keeping logs data
type Logs struct {
	ID             int            `json:"id"`               // [NOT NULL][all] Object identification
	UserID         int            `json:"user_id"`          // [NOT NULL][all] User identification
	Timestamp      time.Time      `json:"timestamp"`        // [NOT NULL][all] Timestamp the logging happened
	Activity       Activity       `json:"activity"`         // [NOT NULL][all] User Activity
	AssignmentId   sql.NullInt64  `json:"assignment_id"`    // [NULLABLE][DeliveredAssignment/FinishedOnePeerReview/PeerReviewDone/CreatAssignment] ID to relative assignment
	CourseID       sql.NullInt64  `json:"course_id"`        // [NULLABLE][JoinedCourse/CreatedCourse] ID to relative course
	SubmissionID   sql.NullInt64  `json:"submission_id"`    // [NULLABLE][DeliveredAssignment/FinishedOnePeerReview/PeerReviewDone] ID to relative submission
	OldValue       sql.NullString `json:"old_value"`        // [NULLABLE][ChangeName/ChangeEmail/ChangeAdminFAQ] Value before changing name/email/faq
	NewValue       sql.NullString `json:"new_value"`        // [NULLABLE][ChangeName/ChangeEmail/ChangeAdminFAQ] Value after changing name/email/faq
	AffectedUserID sql.NullInt64  `json:"affected_user_id"` // [NULLABLE] [FinishedOnePeerReview] Value of student that has had the submission reviewed by one other student
}
