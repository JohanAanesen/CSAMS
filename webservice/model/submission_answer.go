package model

import (
	"database/sql"
	"time"
)

// SubmissionAnswer struct
type SubmissionAnswer struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	AssignmentID int            `json:"assignment_id"`
	SubmissionID int            `json:"submission_id"`
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
	Required     bool           `json:"required"`
}
