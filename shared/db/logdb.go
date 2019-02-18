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

	// Different sql queries to different log types belows
	var rows *sql.Rows
	var err error

	// User changes name or email
	if payload.Activity == model.ChangeEmail || payload.Activity == model.ChangeName {
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `oldvalue`, `newvalue`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.OldValue, payload.NewValue)

		// User changes password
	} else if payload.Activity == model.ChangePassword {
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`) "+
			"VALUES (?, ?)", payload.UserID, payload.Activity)

		// User has delivered assignment, finished peer reviewing or has an assignment that's done with peer-review
	} else if payload.Activity == model.DeliveredAssignment || payload.Activity == model.FinishedPeerReview || payload.Activity == model.PeerReviewDone {
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`,  `submissionid`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID, payload.SubmissionID)

		// Admin has created assignment
	} else if payload.Activity == model.CreatAssignment {
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`) "+
			"VALUES (?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID)

		// User has joined course or admin ahs created course
	} else if payload.Activity == model.JoinedCourse || payload.Activity == model.CreatedCourse {
		rows, err = GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `courseid`) "+
			"VALUES (?, ?, ?)", payload.UserID, payload.Activity, payload.CourseID)

		// Something is wrong
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
