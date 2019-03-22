package model

import (
	"database/sql"
	"time"
)

// ReviewAnswer holds the data for a single review answer
type ReviewAnswer struct {
	ID           int            `json:"id"`
	UserReviewer int            `json:"user_reviewer"`
	UserTarget   int            `json:"user_target"`
	AssignmentID int            `json:"assignment_id"`
	ReviewID     int            `json:"review_id"`
	Type         string         `json:"type"`
	Name         string         `json:"name"`
	Label        string         `json:"label"`
	Description  string         `json:"description"`
	Answer       string         `json:"answer"`
	HasComment   bool           `json:"has_comment"`
	Comment      sql.NullString `json:"comment"`
	Choices      []string       `json:"choices"`
	Submitted    time.Time      `json:"submitted"`
	Weight       int            `json:"weight"`
}
