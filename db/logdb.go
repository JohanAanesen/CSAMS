package db

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"log"
)

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(payload model.Log) bool {

	// Add values in sql query
	rows, err := DB.Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentID`, `courseID`, `userAssignmentID`, `oldvalue`, `newValue`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID, payload.CourseID, payload.UserAssignmentID, payload.OldValue, payload.NewValue)

	// Handle possible error
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	// Close
	defer rows.Close()

	// Nothing went wrong -> return true
	return true
}
