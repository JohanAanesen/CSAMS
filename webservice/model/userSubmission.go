package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
)

// UserSubmission is an struct for user submissions
type UserSubmission struct {
	UserID       int
	AssignmentID int
	SubmissionID int64
	Answers      []Answer2
}

// Answer2 is an struct for one answer
type Answer2 struct {
	Type  string
	Value string
}

// GetUserAnswers returns answers if it exists, empty if not
func GetUserAnswers(userID int, assignmentID int) ([]Answer2, error) {

	// Create an empty answers array
	var answers []Answer2

	// Create query string
	query := "SELECT user_submissions.type, user_submissions.answer FROM user_submissions WHERE user_id =? AND assignment_id=?;"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, userID, assignmentID)
	if err != nil {

		// Returns empty if it fails
		return answers, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		var aType string
		var aValue string

		// Scan rows
		err := rows.Scan(&aType, &aValue)

		// Check for error
		if err != nil {
			return answers, err
		}

		answers = append(answers, Answer2{
			Type:  aType,
			Value: aValue,
		})
	}

	return answers, nil
}

// InsertUserSubmission inserts user submission to the db
func InsertUserSubmission(userSub UserSubmission) error {

	// Go through all answers
	for _, answer := range userSub.Answers {

		// Sql query
		query := "INSERT INTO user_submissions (user_id, submission_id, assignment_id, type, answer) VALUES (?, ?, ?, ?, ?)"
		_, err := db.GetDB().Exec(query, userSub.UserID, userSub.SubmissionID, userSub.AssignmentID, answer.Type, answer.Value)

		// Check if there was an error
		if err != nil {
			return err
		}
	}

	// return nil if no errors
	return nil
}
