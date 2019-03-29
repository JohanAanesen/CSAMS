package repositroy

import "database/sql"

// ReviewAnswerRepository struct
type PeerReviewRepository struct {
	db *sql.DB
}

// NewReviewAnswerRepository func
func NewPeerReviewRepository(db *sql.DB) *PeerReviewRepository {
	return &PeerReviewRepository{
		db: db,
	}
}

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
