package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
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

// FetchAll func
func (repo *LogsRepository) FetchAll() ([]*model.Logs, error) {
	result := make([]*model.Logs, 0)

	// Query to be executed
	query := "SELECT `id`, `userid`, `timestamp`, `activity`, `assignmentid`, `courseid`, `submissionid`, `oldvalue`, `newValue` FROM `logs`"

	// Run query
	rows, err := repo.db.Query(query)
	if err != nil{
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
