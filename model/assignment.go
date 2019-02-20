package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"time"
)

// Assignment hold the data for a single assignment
type Assignment struct {
	ID           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Description  string     `json:"description" db:"description"`
	Created      time.Time  `json:"created" db:"created"`
	Publish      time.Time  `json:"publish" db:"publish"`
	Deadline     time.Time  `json:"deadline" db:"deadline"`
	CourseID     int        `json:"course_id" db:"course_id"`
	SubmissionID int        `json:"-" db:"submission_id"`
	ReviewID     int        `json:"-" db:"review_id"`
	Submission   Submission `json:"submission"`
	Review       Review     `json:"review"`
}

// AssignmentDatabase holds all assignments, and DB-functions
type AssignmentDatabase struct{}

// GetAll returns all assignments in the database
func (adb *AssignmentDatabase) GetAll() []Assignment {
	// TODO (Svein): Implement
	return nil
}

// Insert a new assignment to the database
func (adb *AssignmentDatabase) Insert(a Assignment) (bool, error) {
	rows, err := db.GetDB().Query("...")
	defer rows.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}
