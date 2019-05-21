package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //database driver
	"time"
)

// Assignment hold the data for a single assignment
type Assignment struct {
	ID              int           `json:"id" db:"id"`
	Name            string        `json:"name" db:"name"`
	Description     string        `json:"description" db:"description"`
	Created         time.Time     `json:"created" db:"created"`
	Publish         time.Time     `json:"publish" db:"publish"`
	Deadline        time.Time     `json:"deadline" db:"deadline"`
	CourseID        int           `json:"course_id" db:"course_id"`
	SubmissionID    sql.NullInt64 `json:"-" db:"submission_id"`
	ReviewID        sql.NullInt64 `json:"-" db:"review_id"`
	Submission      Submission    `json:"submission"`
	Review          Review        `json:"review"`
	ReviewEnabled   bool          `json:"review_enabled"`
	ReviewDeadline  time.Time     `json:"review_deadline"`
	Reviewers       sql.NullInt64 `json:"reviewers"`
	ValidationID    sql.NullInt64 `json:"validation_id"`
	GroupDelivery   bool          `json:"group_delivery"`
	MessagesEnabled bool          `json:"messages_enabled"`
}
