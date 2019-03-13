package model

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"log"
)

// activity enum for keeping track of log activity
type activity string

// Enum for logs
const (
	ChangeEmail         activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	FinishedPeerReview  activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCourse        activity = "JOINED-COURSE"                          // User joined course
	CreatedCourse       activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     activity = "ASSIGNMENT-CREATED"                     // Assignment is created
	UpdateAdminFAQ      activity = "UPDATE-ADMIN-FAQ"                       // The admins faq is updated
	// TODO Brede : add more activities later :)
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

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	// Different sql queries to different log types belows
	var rows *sql.Rows
	var err error

	switch payload.Activity {
	case ChangeEmail:
		rows, err = changeEmailUpdateFaq(payload, date)
	case UpdateAdminFAQ:
		rows, err = changeEmailUpdateFaq(payload, date)
	case ChangePassword:
		rows, err = changePassword(payload, date)
	case DeliveredAssignment:
		rows, err = deliveredAssFinishedPeer(payload, date)
	case FinishedPeerReview:
		rows, err = deliveredAssFinishedPeer(payload, date)
	case PeerReviewDone:
		rows, err = deliveredAssFinishedPeer(payload, date)
	case CreatAssignment:
		rows, err = createAssignment(payload, date)
	case JoinedCourse:
		rows, err = joinCreateCourse(payload, date)
	case CreatedCourse:
		rows, err = joinCreateCourse(payload, date)
	default:
		log.Println("Error: Wrong Log.Activity!")
		return false
	}

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

func changeEmailUpdateFaq(data Log, date string) (*sql.Rows, error) {
	rows, err := db.GetDB().Query("INSERT INTO `logs` (`userid`, `timestamp`, `activity`, `oldvalue`, `newvalue`) "+
		"VALUES (?, ?, ?, ?, ?)", data.UserID, date, data.Activity, data.OldValue, data.NewValue)

	return rows, err
}

func changePassword(data Log, date string) (*sql.Rows, error) {
	rows, err := db.GetDB().Query("INSERT INTO `logs` (`userid`, `timestamp`, `activity`) "+
		"VALUES (?, ?, ?)", data.UserID, date, data.Activity)

	return rows, err
}

func deliveredAssFinishedPeer(data Log, date string) (*sql.Rows, error) {
	rows, err := db.GetDB().Query("INSERT INTO `logs` (`userid`, `timestamp`,  `activity`, `assignmentid`,  `submissionid`) "+
		"VALUES (?, ?, ?, ?, ?)", data.UserID, date, data.Activity, data.AssignmentID, data.SubmissionID)

	return rows, err
}

func createAssignment(data Log, date string) (*sql.Rows, error) {
	rows, err := db.GetDB().Query("INSERT INTO `logs` (`userid`, `timestamp`, `activity`, `assignmentid`) "+
		"VALUES (?, ?, ?, ?)", data.UserID, date, data.Activity, data.AssignmentID)

	return rows, err
}

func joinCreateCourse(data Log, date string) (*sql.Rows, error) {
	rows, err := db.GetDB().Query("INSERT INTO `logs` (`userid`, `timestamp`,  `activity`, `courseid`) "+
		"VALUES (?, ?, ?, ?)", data.UserID, date, data.Activity, data.CourseID)

	return rows, err
}
