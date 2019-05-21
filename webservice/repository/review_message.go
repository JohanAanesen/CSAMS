package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

// UserRepository struct
type ReviewMessageRepository struct {
	db *sql.DB
}

// NewUserRepository return a pointer to a new UserRepository
func NewReviewMessageRepository(db *sql.DB) *ReviewMessageRepository {
	return &ReviewMessageRepository{
		db: db,
	}
}

// Insert message function
func (repo *ReviewMessageRepository) Insert(message model.ReviewMessage) error {

	query := "INSERT INTO review_messages (user_reviewer, user_target, assignment_id, message) VALUES (?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, message.UserReviewer, message.UserTarget, message.AssignmentID, message.Message)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return err
}

// FetchAllForAssignmentUser messages function
func (repo *ReviewMessageRepository) FetchAllForAssignmentUser(assignmentID int, userID int) ([]*model.ReviewMessage, error) {
	result := make([]*model.ReviewMessage, 0)

	query := "SELECT id, user_reviewer, user_target, assignment_id, message FROM review_messages WHERE assignment_id = ? AND user_target = ?"

	rows, err := repo.db.Query(query, assignmentID, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewMessage{}

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.AssignmentID, &temp.Message)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}
