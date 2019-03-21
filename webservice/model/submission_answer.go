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
	HasComment   bool           `json:"has_comment"`
	Answer       string         `json:"answer"`
	Comment      sql.NullString `json:"comment"`
	Submitted    time.Time      `json:"submitted"`
}
