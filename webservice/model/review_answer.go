package model

import "time"

// ReviewAnswer holds the data for a single review answer
type ReviewAnswer struct {
	ID           int       `json:"id"`
	UserReviewer int       `json:"user_reviewer"`
	UserTarget   int       `json:"user_target"`
	ReviewID     int       `json:"review_id"`
	AssignmentID int       `json:"assignment_id"`
	Type         string    `json:"type"`
	Name         string    `json:"name"`
	Label        string    `json:"label"`
	Answer       string    `json:"answer"`
	Comment      string    `json:"comment"`
	Submitted    time.Time `json:"submitted"`
}