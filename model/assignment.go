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
func (repo *AssignmentRepository) Insert(assignment Assignment) error {
	// Create query string
	query := "INSERT INTO assignments (name, description, publish, deadline, course_id) VALUES (?, ?, ?, ?, ?);"
	// Prepare and execute query
	rows, err := db.GetDB().Exec(query, assignment.Name, assignment.Description, assignment.Publish, assignment.Deadline, assignment.CourseID)
	// Check for error
	if err != nil {
		return err
	}

	// Get last inserted ID
	id, err := rows.LastInsertId()
	// Check for error
	if err != nil {
		return err
	}

	// Check if we have set a submission_id
	if assignment.SubmissionID != 0 {
		// Create query string
		query := "UPDATE assignments SET submission_id = ? WHERE id = ?;"
		// Prepare and execute query
		rows, err := db.GetDB().Query(query, assignment.SubmissionID, id)
		// Check for error
		if err != nil {
			return err
		}
		// Close connection
		defer rows.Close()
	}

	// Check if we have set a review_id
	if assignment.ReviewID != 0 {
		// Create query string
		query := "UPDATE assignments SET review_id = ? WHERE id = ?;"
		// Prepare and execute query
		rows, err := db.GetDB().Query(query, assignment.ReviewID, id)
		// Check for error
		if err != nil {
			return err
		}
		// Close connection
		defer rows.Close()
	}


	return nil
}
