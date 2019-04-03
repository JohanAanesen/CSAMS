package model

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"log"
)

// Activity enum for keeping track of log Activity
type Activity string

// Enum for logs
const (
	ChangeEmail         Activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      Activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment Activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	UpdateAssignment    Activity = "ASSIGNMENT-UPDATED"                     // User re-delivered assignment
	DeleteAssignment    Activity = "ASSIGNMENT-DELETE"                      // user deleted assignment
	FinishedPeerReview  Activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      Activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCourse        Activity = "JOINED-COURSE"                          // User joined course
	CreatedCourse       Activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     Activity = "ASSIGNMENT-CREATED"                     // Assignment is created
	UpdateAdminFAQ      Activity = "UPDATE-ADMIN-FAQ"                       // The admins faq is updated
	NewUser             Activity = "NEW-USER"                               // A new user is created
	// TODO Brede : add more activities later :)
)

// Log struct to hold log-data
type Log struct {
	UserID       int      // [NOT NULL][all] User identification
	Activity     Activity // [NOT NULL][all] User Activity
	IsTeacher    bool     // [NULLABLE][later user] Says if the user is student or teacher (This is later checked from database)
	AssignmentID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone/CreatAssignment] ID to relative assignment
	CourseID     int      // [NULLABLE][JoinedCourse/CreatedCourse] ID to relative course
	SubmissionID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone] ID to relative submission
	OldValue     string   // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value before changing name/email/faq
	NewValue     string   // [NULLABLE][ChangeName/ChangeEmail/UpdateAdminFAQ] Value after changing name/email/faq
}

/* TODO : convert to something like this later
func LogChangeEmail(userID int, oldValue string, newValue string) error {

}

func LogChangeFAQ(userID int, oldValue string, newValue string) error {

}

func LogCreateAssignment(userID int, assignmentID string) error {

}

func LogDeliveredAssignment(userID int, assignmentID string, submissionID string) error {

}

func LogUpdateAssignment(userID int, assignmentID string, submissionID string) error {

}

func LogDeleteAssignment(userID int, assignmentID string, submissionID string) error {

}

func LogPeerReviewFinished(userID int, assignmentID string, submissionID string) error {

}

func LogPeerReviewDone(userID int, assignmentID string, submissionID string) error {

}

func LogCreateCourse(userID int, courseID string) error {

}

func LogJoinCourse(userID int, courseID string) error {

}
*/

// LogToDB adds logs to the database when an user/admin does something noteworthy
func LogToDB(payload Log) error {

	// UserID and Activity can not be nil
	if payload.UserID <= 0 || payload.Activity == "" {
		return errors.New("error: userid and/or Activity can not be nil (log.db)")
	}

	// Different sql queries to different log types belows
	var err error

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		return err
	}

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	switch payload.Activity {
	case ChangeEmail:
		err = changeEmailUpdateFaq(tx, payload, date)
	case UpdateAdminFAQ:
		err = changeEmailUpdateFaq(tx, payload, date)
	case ChangePassword:
		err = changePassword(tx, payload, date)
	case DeliveredAssignment:
		err = deliveredAssFinishedPeer(tx, payload, date)
	case UpdateAssignment:
		err = deliveredAssFinishedPeer(tx, payload, date)
	case DeleteAssignment:
		err = deliveredAssFinishedPeer(tx, payload, date)
	case FinishedPeerReview:
		err = deliveredAssFinishedPeer(tx, payload, date)
	case PeerReviewDone:
		err = deliveredAssFinishedPeer(tx, payload, date)
	case CreatAssignment:
		err = createAssignment(tx, payload, date)
	case JoinedCourse:
		err = joinCreateCourse(tx, payload, date)
	case CreatedCourse:
		err = joinCreateCourse(tx, payload, date)
	case NewUser:
		err = newUser(tx, payload, date)
	default:
		log.Println("Error: Wrong Log.Activity!")
		return errors.New("error: wrong log.Activity type (log.db)")
	}

	// Handle possible error
	if err != nil {
		tx.Rollback() //quit transaction if error
		return err
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		return err
	}

	// Nothing went wrong -> return true
	return nil
}

func changeEmailUpdateFaq(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`, `oldvalue`, `newvalue`) "+
		"VALUES (?, ?, ?, ?, ?)", data.UserID, date, data.Activity, data.OldValue, data.NewValue)

	return err
}

func changePassword(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", data.UserID, date, data.Activity)

	return err
}

func deliveredAssFinishedPeer(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`,  `Activity`, `assignmentid`,  `submissionid`) "+
		"VALUES (?, ?, ?, ?, ?)", data.UserID, date, data.Activity, data.AssignmentID, data.SubmissionID)

	return err
}

func createAssignment(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`, `assignmentid`) "+
		"VALUES (?, ?, ?, ?)", data.UserID, date, data.Activity, data.AssignmentID)

	return err
}

func joinCreateCourse(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`,  `Activity`, `courseid`) "+
		"VALUES (?, ?, ?, ?)", data.UserID, date, data.Activity, data.CourseID)

	return err
}

func newUser(tx *sql.Tx, data Log, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", data.UserID, date, data.Activity)

	return err
}
