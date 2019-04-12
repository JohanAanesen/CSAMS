package model

import (
	"database/sql"
	"time"
)

// Activity enum for keeping track of log Activity
type Activity string

// Enum for logs
const (
	// System logs
	NewUser             Activity = "NEW-USER"                      // A new user is created
	ChangeEmail         Activity = "CHANGE-EMAIL"                  // User changed email
	ChangePassword      Activity = "CHANGE-PASSWORD"               // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	ChangePasswordEmail Activity = "CHANGE-PASSWORD-THROUGH-EMAIL" // User changed password through email

	// Course logs
	DeliveredAssignment   Activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	UpdateAssignment      Activity = "ASSIGNMENT-UPDATED"                     // User re-delivered assignment
	FinishedOnePeerReview Activity = "FINISHED-ONE-PEER-REVIEW"               // User is done with one peer review (that this user did)
	ChangeOnePeerReview   Activity = "CHANGE-ONE-PEER-REVIEW"                 // User changed one peer review
	PeerReviewDone        Activity = "PEER-REVIEW-IS-DONE-FOR-ONE-SUBMISSION" // User submission is finished peer-reviewed by other students
	JoinedCourse          Activity = "JOINED-COURSE"                          // User joined course

	// Admin logs
	DeleteAssignment           Activity = "ASSIGNMENT-DELETE"                // user deleted assignment
	CreatAssignment            Activity = "ASSIGNMENT-CREATED"               // Assignment is created
	CreatedCourse              Activity = "COURSE-CREATED"                   // Course is created
	AdminChangeFaq             Activity = "ADMIN-CHANGE-FAQ"                 // Admin changed the FAQ
	AdminEmailCourseStudents   Activity = "ADMIN-EMAILED-STUDENTS-IN-COURSE" // Admin emailed all students in course through the system
	AdminRemoveUserFromCOurse  Activity = "ADMIN-REMOVED-USER-FROM-COURSE"   // Admin removed one user from course
	AdminChangeSTudentPassword Activity = "ADMIN-CHANGE-STUDENT-PASSWORD"    // Admin changed one users password

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
	OldValue       sql.NullString `json:"old_value"`        // [NULLABLE][ChangeName/ChangeEmail/AdminChangeFaq] Value before changing name/email/faq
	NewValue       sql.NullString `json:"new_value"`        // [NULLABLE][ChangeName/ChangeEmail/AdminChangeFaq] Value after changing name/email/faq
	AffectedUserID sql.NullInt64  `json:"affected_user_id"` // [NULLABLE] [FinishedOnePeerReview] Value of student that has had the submission reviewed by one other student
}
