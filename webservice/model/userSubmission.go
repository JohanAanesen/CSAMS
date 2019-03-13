package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"time"
)

// UserSubmission is an struct for user submissions
type UserSubmission struct {
	UserID       int
	AssignmentID int
	SubmissionID int64
	Answers      []Answer
	Submitted    time.Time
}

// GetUserAnswers returns answers if it exists, empty if not
// TODO (Svein): Move this into some struct as a method, or rename it to reflect it's actions
func GetUserAnswers(userID int, assignmentID int) ([]Answer, error) {
	// Create an empty answers array
	var result []Answer
	// Create query string
	query := "SELECT id, type, answer, comment FROM user_submissions WHERE user_id =? AND assignment_id=?;"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, userID, assignmentID)
	if err != nil {

		// Returns empty if it fails
		return result, err
	}
	// Close connection
	defer rows.Close()
	// Loop through results
	for rows.Next() {
		var temp Answer
		// Scan rows
		err := rows.Scan(&temp.ID, &temp.Type, &temp.Value, &temp.Comment)
		// Check for error
		if err != nil {
			return result, err
		}
		result = append(result, temp)
	}

	return result, nil
}

// GetSubmittedTime returns submitted time if it exists, empty if not
func GetSubmittedTime(userID int, assignmentID int) (time.Time, bool, error) {
	result := time.Time{}

	// Create query string
	query := "SELECT DISTINCT submitted FROM user_submissions WHERE user_id=? AND assignment_id=?;"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, userID, assignmentID)
	if err != nil {

		// Returns empty if it fails
		return result, false, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	if rows.Next() {
		// Scan rows
		err := rows.Scan(&result)
		// Check for error
		if err != nil {
			return time.Time{}, false, err
		}

		return result, true, nil
	}

	return time.Time{}, false, nil
}

// UploadUserSubmission uploads user submission to the db
func UploadUserSubmission(userSub UserSubmission) error {
	// Query string
	query := "INSERT INTO user_submissions (user_id, submission_id, assignment_id, type, answer, comment) " +
		"VALUES (?, ?, ?, ?, ?, ?)"
	// Begin transaction with database
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	// Go through all answers
	for _, answer := range userSub.Answers {
		// Sql query
		_, err := db.GetDB().Exec(query, userSub.UserID, userSub.SubmissionID, userSub.AssignmentID,
			answer.Type, answer.Value, answer.Comment)
		// Check if there was an error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// return nil if no errors
	return err
}

// UpdateUserSubmission updates user submission to the db
func UpdateUserSubmission(userSub UserSubmission) error {
	// Sql query
	query := "UPDATE user_submissions SET answer=?, submitted=?, comment=? WHERE id=?"
	// Norwegian time TODO time
	now := time.Now().UTC().Add(time.Hour)

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	// Go through all answers
	for _, answer := range userSub.Answers {
		_, err := db.GetDB().Exec(query, answer.Value, now, answer.Comment, answer.ID)
		// Check if there was an error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// return nil if no errors
	return nil
}
