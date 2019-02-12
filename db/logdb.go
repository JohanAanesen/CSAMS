package db

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"log"
)

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(payload model.Log) bool {

	// UserID and Activity can not be nil
	if payload.UserID <= 0 || payload.Activity == "" {
		return false
	}

	// Add values in sql query
	rows, err := DB.Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`, `courseid`, `submissionid`, `oldvalue`, `newvalue`) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?)", payload.UserID, payload.Activity, nil, nil, nil, payload.OldValue, payload.NewValue)
	// TODO : remove nil and fix it

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

func ReadAllLogs() {
	// TODO : add code for getting all logs in nice format here
}
