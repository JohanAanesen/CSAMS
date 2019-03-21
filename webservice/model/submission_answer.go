package model

import "time"

// SubmissionAnswer struct
type SubmissionAnswer struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	AssignmentID int       `json:"assignment_id"`
	SubmissionID int       `json:"submission_id"`
	Type         string    `json:"type"`
	Answer       string    `json:"answer"`
	Comment      string    `json:"comment"`
	Submitted    time.Time `json:"submitted"`
}
