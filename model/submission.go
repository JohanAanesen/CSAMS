package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"strings"
)

// Submission .... TODO (Svein): comment
type Submission struct {
	ID     int  `json:"id" db:"id"`
	FormID int  `json:"-" db:"form_id"`
	Form   Form `json:"form"`
}

// SubmissionRepository ... TODO (Svein): comment
type SubmissionRepository struct {
}

// Insert form and fields to database
func (repo *SubmissionRepository) Insert(form Form) error {
	// Declare FormRepository
	var formRepo = FormRepository{}

	// Insert form to database, and get last inserted Id from it
	formID, err := formRepo.Insert(form)
	if err != nil {
		return err
	}

	// Insertions query
	query := "INSERT INTO submissions (form_id) VALUES(?)"
	// Insert form_id into submissions
	rows, err := db.GetDB().Query(query, formID)
	// Check for error
	if err != nil {
		return err
	}
	// Close the connections
	defer rows.Close()

	// Loop trough fields in the forms
	for _, field := range form.Fields {
		// Join the array to a single string for 'choices'
		choices := strings.Join(field.Choices, ",")
		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		// Execute the query
		rows, err := db.GetDB().Query(query, int(formID), field.Type, field.Name, field.Label, field.Description, field.Order, field.Weight, choices)
		// Check for error
		if err != nil {
			return err
		}
		// Close the connection
		rows.Close()
	}

	// Return no error
	return nil
}

// GetAll returns all submission in the database
func (repo *SubmissionRepository) GetAll() ([]Submission, error) {
	// Declare return slice
	var result []Submission
	// Create query-string
	query := "SELECT * FROM submissions"
	// Perform query
	rows, err := db.GetDB().Query(query)
	// Check for error
	if err != nil {
		return []Submission{}, err
	}
	// Close connection
	defer rows.Close()

	// Loop through rows
	for rows.Next() {
		// Declare a single Submission
		var submission = Submission{}
		// Scan the data from the rows
		err = rows.Scan(&submission.ID, &submission.FormID)
		// Check for error
		if err != nil {
			return []Submission{}, err
		}

		// Append scan-result to result
		result = append(result, submission)
	}

	// Declare a FormRepository
	var formRepo = FormRepository{}
	// Loop through all submissions
	for index, submission := range result {
		// Get form from database
		form, err := formRepo.Get(submission.FormID)
		// Check for error
		if err != nil {
			return []Submission{}, nil
		}
		// Get the form
		submission.Form = form
		// Set the new values
		result[index] = submission
	}

	return result, nil
}
