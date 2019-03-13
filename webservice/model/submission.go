package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
)

// Submission struct
type Submission struct {
	ID     int  `json:"id" db:"id"`
	FormID int  `json:"-" db:"form_id"`
	Form   Form `json:"form"`
}

// SubmissionRepository struct
type SubmissionRepository struct{}

// Insert form and fields to database
func (repo *SubmissionRepository) Insert(form Form) error {
	// Declare FormRepository
	var formRepo = FormRepository{}

	// Insert form to database, and get last inserted Id from it
	formID, err := formRepo.Insert(form)
	if err != nil {
		return err
	}

	//Start transaction
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	// Insertions query
	query := "INSERT INTO submissions (form_id) VALUES(?)"
	// Insert form_id into submissions
	_, err = tx.Exec(query, formID)
	// Check for error
	if err != nil {
		tx.Rollback() //rollback if err
		return err
	}

	// Loop trough fields in the forms
	for _, field := range form.Fields {

		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices, hasComment) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"

		var hasComment = 0
		if field.HasComment {
			hasComment = 1
		}

		//Execute query
		_, err = tx.Exec(query, formID, field.Type, field.Name, field.Label, field.Description,
			field.Order, field.Weight, field.Choices, hasComment)
		if err != nil {
			tx.Rollback() //rollback if err
			return err
		}
	}

	err = tx.Commit() //commit transaction/changes
	if err != nil {
		return err
	}

	// Return no error
	return nil
}

// GetAll returns all submission in the database
func (repo *SubmissionRepository) GetAll() ([]Submission, error) {
	// Declare return slice
	var result []Submission
	// Create query-string
	query := "SELECT id, form_id FROM submissions"
	// Perform query
	rows, err := db.GetDB().Query(query)
	// Check for error
	if err != nil {
		return result, err
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

// DeleteSubmissionsToAssignment deletes specific submissions to assignmentID and SubmissionID
func (repo *SubmissionRepository) DeleteSubmissionsToAssignment(assID int, subID int64) error {
	// Create query-string
	query := "DELETE  FROM user_submissions WHERE assignment_id = ? AND submission_id = ?"
	// Perform query
	rows, err := db.GetDB().Query(query, assID, subID)
	// Check for error
	if err != nil {
		return err
	}
	// Close connection
	defer rows.Close()

	return nil
}

// GetSubmissionsCountFromAssignment returns all submission in the database from specific assignment and submission form
func (repo *SubmissionRepository) GetSubmissionsCountFromAssignment(assID int, subID int64) (int, error) {
	// Declare return slice
	var result int
	// Create query-string
	query := "select count(distinct user_id) from user_submissions WHERE assignment_id LIKE ? AND submission_id LIKE ?"
	// Perform query
	rows, err := db.GetDB().Query(query, assID, subID)
	// Check for error
	if err != nil {
		return 0, err
	}
	// Close connection
	defer rows.Close()

	// Loop through rows
	if rows.Next() {
		// Scan the data from the row
		err = rows.Scan(&result)
		// Check for error
		if err != nil {
			return 0, err
		}

	}

	return result, nil
}

// Update a form in the database
// Deletes all fields, and recreates them
func (repo *SubmissionRepository) Update(form Form) error {
	query := "UPDATE forms SET prefix=?, name=? WHERE id=?"
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	_, err = db.GetDB().Exec(query, form.Prefix, form.Name, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "DELETE FROM fields WHERE form_id=?"
	_, err = db.GetDB().Exec(query, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Loop trough fields in the forms
	for _, field := range form.Fields {
		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices, hasComment) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"

		var hasComment = 0
		if field.HasComment {
			hasComment = 1
		}
		// Execute the query
		_, err := db.GetDB().Exec(query, form.ID, field.Type, field.Name, field.Label, field.Description,
			field.Order, field.Weight, field.Choices, hasComment)
		// Check for error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	// Return no error
	return nil
}

// Delete a review form based on it's id
func (repo *SubmissionRepository) Delete(id int) error {
	// SQL query
	query := "DELETE FROM fields WHERE form_id=?"
	// Begin transaction
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	// Execute query
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	// SQL query
	query = "DELETE FROM submissions WHERE form_id=?"
	// Begin transaction
	tx, err = db.GetDB().Begin()
	if err != nil {
		return err
	}
	// Execute transaction
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	// SQL query
	query = "DELETE FROM forms WHERE id=?"
	// Begin transaction
	tx, err = db.GetDB().Begin()
	if err != nil {
		return err
	}
	// Execute transaction
	_, err = tx.Exec(query, id)
	if err != nil {
		// Rollback transaction on error
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		// Rollback transaction on error
		tx.Rollback()
		return err
	}

	return err
}
