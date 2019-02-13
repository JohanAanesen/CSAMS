package db

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"log"
)

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(payload model.Log) bool {

	// UserID and Activity can not be nil
	if payload.UserID <= 0 || payload.Activity == "" {
		return false
	}

	// TODO : Refactor the shit out of this function
	var rows *sql.Rows
	var err error

	if payload.Activity == model.ChangeEmail || payload.Activity == model.ChangeName {
		// Add values in sql query
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `oldvalue`, `newvalue`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.OldValue, payload.NewValue)

	} else if payload.Activity == model.ChangePassword {
		// Add values in sql query
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`) "+
			"VALUES (?, ?)", payload.UserID, payload.Activity)

	} else if payload.Activity == model.DeliveredAssignment || payload.Activity == model.FinishedPeerReview || payload.Activity == model.PeerReviewDone || payload.Activity == model.CreatAssignment {
		// Add values in sql query
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`,  `submissionid`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID, payload.SubmissionID)

	} else if payload.Activity == model.JoinedCousrse || payload.Activity == model.CreatedCourse {
		// Add values in sql query
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `courseid`) "+
			"VALUES (?, ?, ?)", payload.UserID, payload.Activity, payload.CourseID)
		
	} else {
		return false
	}
	// TODO ends here

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
