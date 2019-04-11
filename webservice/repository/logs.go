package repository

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/CSAMS/schedulerservice/db"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"log"
)

// LogsRepository struct
type LogsRepository struct {
	db *sql.DB
}

// NewLogsRepository func
func NewLogsRepository(db *sql.DB) *LogsRepository {
	return &LogsRepository{
		db: db,
	}
}

// FetchAll fetches all logs
func (repo *LogsRepository) FetchAll() ([]*model.Logs, error) {
	result := make([]*model.Logs, 0)

	// Query to be executed
	query := "SELECT `id`, `userid`, `timestamp`, `activity`, `assignmentid`, `courseid`, `submissionid`, `oldvalue`, `newValue` FROM `logs`"

	// Run query
	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	// Close rows
	defer rows.Close()

	// Go through all rows
	for rows.Next() {
		temp := model.Logs{}

		// Add to temporary struct
		err = rows.Scan(&temp.ID, &temp.UserID, &temp.Timestamp, &temp.Activity,
			&temp.AssignmentId, &temp.CourseID, &temp.SubmissionID, &temp.OldValue,
			&temp.NewValue)
		if err != nil {
			return result, err
		}

		// Append to result array
		result = append(result, &temp)
	}

	return result, err
}

// Insert inserts all types of logs
func (repo *LogsRepository) Insert(logx model.Logs) error {

	// Different sql queries to different log types belows
	var err error

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		return err
	}

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	// Check what needs to be logged
	switch logx.Activity {
	case model.ChangeEmail:
		//err = changeEmailUpdateFaq(tx, logx, date)
	case model.UpdateAdminFAQ:
		//err = changeEmailUpdateFaq(tx, logx, date)
	case model.ChangePassword:
		//err = changePassword(tx, logx, date)
	case model.DeliveredAssignment:
		//err = deliveredAssFinishedPeer(tx, logx, date)
	case model.UpdateAssignment:
		//err = deliveredAssFinishedPeer(tx, logx, date)
	case model.DeleteAssignment:
		//err = deliveredAssFinishedPeer(tx, logx, date)
	case model.FinishedPeerReview:
		//err = deliveredAssFinishedPeer(tx, logx, date)
	case model.PeerReviewDone:
		//err = deliveredAssFinishedPeer(tx, logx, date)
	case model.CreatAssignment:
		//err = createAssignment(tx, logx, date)
	case model.JoinedCourse:
		//err = joinCreateCourse(tx, logx, date)
	case model.CreatedCourse:
		//err = joinCreateCourse(tx, logx, date)
	case model.NewUser:
		err = newUser(tx, logx, date)
	default:
		log.Println("error: wrong log.activity!")
		return errors.New("error: wrong log.activity type")
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

	// Nothing went wrong -> return nil errors
	return nil
}

/* TODO brede add this
func changeEmailUpdateFaq(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`, `oldvalue`, `newvalue`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.OldValue, logx.NewValue)

	return err
}

func changePassword(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", logx.UserID, date, logx.Activity)

	return err
}

func deliveredAssFinishedPeer(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`,  `Activity`, `assignmentid`,  `submissionid`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentId, logx.SubmissionID)

	return err
}

func createAssignment(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`, `assignmentid`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentId)

	return err
}

func joinCreateCourse(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`,  `Activity`, `courseid`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.CourseID)

	return err
}
*/

// newUser query for inserting new user log
func newUser(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`userid`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", logx.UserID, date, logx.Activity)

	return err
}
