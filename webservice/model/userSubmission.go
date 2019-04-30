package model

import (
	"time"
)

// UserSubmission is an struct for user submissions
type UserSubmission struct {
	UserID       int
	AssignmentID int
	SubmissionID int64
	Answers      []Answer
	Submitted    time.Time
}
