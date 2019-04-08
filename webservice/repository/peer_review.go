package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

// PeerReviewRepository struct
type PeerReviewRepository struct {
	db *sql.DB
}

// NewPeerReviewRepository func
func NewPeerReviewRepository(db *sql.DB) *PeerReviewRepository {
	return &PeerReviewRepository{
		db: db,
	}
}

// TargetExists Checks if the target exist in the table
func (repo *PeerReviewRepository) TargetExists(assignmentID int, userID int) (bool, error) {
	var result int

	query := "SELECT COUNT(DISTINCT user_id) FROM peer_reviews WHERE user_id = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, userID, assignmentID)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return false, err
		}

		// If the query found the user
		if result == 1 {
			return true, nil
		}
	}

	return false, err
}

// FetchPeerReviewsFromAssignment func
func (repo *PeerReviewRepository) FetchPeerReviewsFromAssignment(assignmentID int) ([]*model.PeerReview, error) {
	result := make([]*model.PeerReview, 0)
	query := "SELECT pr.id, pr.user_id, u2.name, pr.review_user_id, u.name, pr.assignment_id FROM peer_reviews AS pr INNER JOIN users AS u ON pr.review_user_id = u.id INNER JOIN users AS u2 ON pr.user_id = u2.id WHERE pr.assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var temp model.PeerReview

		err = rows.Scan(&temp.ID, &temp.ReviewerID, &temp.ReviewerName, &temp.TargetID, &temp.TargetName, &temp.AssignmentID)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}