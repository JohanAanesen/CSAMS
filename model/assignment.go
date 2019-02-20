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

// AssignmentRepository holds all assignments, and DB-functions
type AssignmentRepository struct{}

// GetAll returns all assignments in the database
func (repo *AssignmentRepository) GetAll() []Assignment {
	// TODO (Svein): Implement
	return nil
}

// Insert a new assignment to the database
func (repo *AssignmentRepository) Insert(assignment Assignment) (bool, error) {
	query := "INSERT INTO assignments (name, description, publish, deadline, course_id, submission_id, review_id) VALUES (?, ?, ?, ?, ?, ?, ?);"
	rows, err := db.GetDB().Query(query, assignment.Name, assignment.Description, assignment.Publish, assignment.Deadline, assignment.CourseID, assignment.SubmissionID, assignment.ReviewID)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return true, nil
}
