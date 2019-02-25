package model

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"log"
)

// activity enum for keeping track of log activity
type activity string

// Enum for logs
const (
	ChangeName          activity = "CHANGE-NAME"                            // User changed name
	ChangeEmail         activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	FinishedPeerReview  activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCourse        activity = "JOINED-COURSE"                          // User joined course
	CreatedCourse       activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     activity = "ASSIGNMENT-CREATED"                     // Assignment is created
	UpdateAdminFAQ      activity = "UPDATE-ADMIN-FAQ"                       // The admins faq is updated
)

// Log struct to hold log-data
type Log struct {
	UserID       int      // [NOT NULL][all] User identification
	Activity     activity // [NOT NULL][all] User activity
	IsTeacher    bool     // [NULLABLE][later user] Says if the user is student or teacher (This is later checked from database)
	AssignmentID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone/CreatAssignment] ID to relative assignment
	CourseID     int      // [NULLABLE][JoinedCourse/CreatedCourse] ID to relative course
	SubmissionID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone] ID to relative submission
	OldValue     string   // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value before changing name/email/faq
	NewValue     string   // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value after changing name/email/faq
}

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(payload Log) bool {

	// UserID and Activity can not be nil
	if payload.UserID <= 0 || payload.Activity == "" {
		return false
	}

	// TODO : Refactor the shit out of this function

	// Different sql queries to different log types belows
	var rows *sql.Rows
	var err error

	// User changes name or email
	if payload.Activity == ChangeEmail || payload.Activity == ChangeName || payload.Activity == UpdateAdminFAQ {
		rows, err = db.GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `oldvalue`, `newvalue`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.OldValue, payload.NewValue)

		// User changes password
	} else if payload.Activity == ChangePassword {
		rows, err = db.GetDB().Query("INSERT INTO `logs` (`userid`, `activity`) "+
			"VALUES (?, ?)", payload.UserID, payload.Activity)

		// User has delivered assignment, finished peer reviewing or has an assignment that's done with peer-review
	} else if payload.Activity == DeliveredAssignment || payload.Activity == FinishedPeerReview || payload.Activity == PeerReviewDone {
		rows, err = db.GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`,  `submissionid`) "+
			"VALUES (?, ?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID, payload.SubmissionID)

		// Admin has created assignment
	} else if payload.Activity == CreatAssignment {
		rows, err = db.GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `assignmentid`) "+
			"VALUES (?, ?, ?)", payload.UserID, payload.Activity, payload.AssignmentID)

		// User has joined course or admin ahs created course
	} else if payload.Activity == JoinedCourse || payload.Activity == CreatedCourse {
		rows, err = db.GetDB().Query("INSERT INTO `logs` (`userid`, `activity`, `courseid`) "+
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
