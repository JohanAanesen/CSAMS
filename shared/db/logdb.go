package db

import (
	"log"
	"time"
)

type category string

const (
	ChangeName     category = "CHANGE-NAME"
	ChangeEmail    category = "CHANGE-EMAIL"
	ChangePassword category = "CHANGE-PASSWORD"
)

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(userID int, category category) bool {

	timeStamp := time.Now().UnixNano()

	// Convert back to date
	// tm := time.Unix(0, timeStamp)

	// Add values in sql query
	rows, err := GetDB().Query("INSERT INTO logs(userid, timestamp, log) VALUES (?, ?, ?)", userID, timeStamp, category)

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
