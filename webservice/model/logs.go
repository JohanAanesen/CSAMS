package model

import (
	"database/sql"
	"time"
)

// Activity enum for keeping track of log Activity
type Activity string

// Enum for logs
const (
	ChangeEmail         Activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      Activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment Activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	UpdateAssignment    Activity = "ASSIGNMENT-UPDATED"                     // User re-delivered assignment
	DeleteAssignment    Activity = "ASSIGNMENT-DELETE"                      // user deleted assignment
	FinishedPeerReview  Activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      Activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCourse        Activity = "JOINED-COURSE"                          // User joined course
	CreatedCourse       Activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     Activity = "ASSIGNMENT-CREATED"                     // Assignment is created
	UpdateAdminFAQ      Activity = "UPDATE-ADMIN-FAQ"                       // The admins faq is updated
	NewUser             Activity = "NEW-USER"                               // A new user is created
	// TODO Brede : add more activities later :)
)

// Logs struct for keeping logs data
type Logs struct {
	ID           int            `json:"id"`            // [NOT NULL][all] Object identification
	UserID       int            `json:"user_id"`       // [NOT NULL][all] User identification
	Timestamp    time.Time      `json:"timestamp"`     // [NOT NULL][all] Timestamp the logging happened
	Activity     Activity       `json:"activity"`      // [NOT NULL][all] User Activity
	AssignmentId sql.NullInt64  `json:"assignment_id"` // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone/CreatAssignment] ID to relative assignment
	CourseID     sql.NullInt64  `json:"course_id"`     // [NULLABLE][JoinedCourse/CreatedCourse] ID to relative course
	SubmissionID sql.NullInt64  `json:"submission_id"` // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone] ID to relative submission
	OldValue     sql.NullString `json:"old_value"`     // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value before changing name/email/faq
	NewValue     sql.NullString `json:"new_value"`     // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value after changing name/email/faq
}