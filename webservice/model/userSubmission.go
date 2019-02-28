package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
)

// UserSubmission is an struct for user submissions
type UserSubmission struct {
	UserID       int
	SubmissionID int
	Answers      []Answer
}

// Answer is an struct for one answer
type Answer struct {
	Type  string
	Value string
}

/*
// GetUserSubmission returns userSubmission if it exists
func GetUserSubmission(userID int, assignmentID int) (UserSubmission, error) {

	// Declare empty slice
	var result []Assignment

	// Create query string
	query := "SELECT id, name, description, created, publish, deadline, course_id FROM assignments;"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query)
	if err != nil {
		return nil, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		// Declare empty struct
		var assignment Assignment
		// Scan rows
		err := rows.Scan(&assignment.ID, &assignment.Name, &assignment.Description,
			&assignment.Created, &assignment.Publish, &assignment.Deadline,
			&assignment.CourseID)
		// Check for error
		if err != nil {
			return nil, err
		}

		// Append retrieved row
		result = append(result, assignment)
	}

	return result, nil
}
*/
// InsertUserSubmission inserts user submission to the db
func InsertUserSubmission(userSub UserSubmission) error {

	// Go through all answers
	for _, answer := range userSub.Answers {

		// Sql query
		query := "INSERT INTO user_submissions (user_id, submission_id, field_name, answer) VALUES (?, ?, ?, ?)"
		_, err := db.GetDB().Exec(query, userSub.UserID, userSub.SubmissionID, answer.Type, answer.Value)

		// Check if there was an error
		if err != nil {
			return err
		}
	}

	// return nil if no errors
	return nil
}
