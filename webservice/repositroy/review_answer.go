package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
)

// ReviewAnswerRepository struct
type ReviewAnswerRepository struct {
	db *sql.DB
}

// NewReviewAnswerRepository func
func NewReviewAnswerRepository(db *sql.DB) *ReviewAnswerRepository {
	return &ReviewAnswerRepository{
		db: db,
	}
}

// FetchForAssignment func
func (repo *ReviewAnswerRepository) FetchForAssignment(assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT * FROM user_reviews WHERE assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		temp := model.ReviewAnswer{}

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Answer, &temp.Comment, &temp.Submitted)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchForTarget func
func (repo *ReviewAnswerRepository) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT * FROM user_reviews WHERE user_target = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, target, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		temp := model.ReviewAnswer{}

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Answer, &temp.Comment, &temp.Submitted)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchForReviewer func
func (repo *ReviewAnswerRepository) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted FROM user_reviews WHERE user_reviewer = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, reviewer, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		temp := model.ReviewAnswer{}

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Answer, &temp.Comment, &temp.Submitted)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchForReviewerAndTarget func
func (repo *ReviewAnswerRepository) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT id, user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted FROM user_reviews WHERE user_reviewer = ? AND user_target = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, reviewer, target, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		temp := model.ReviewAnswer{}

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Answer, &temp.Comment, &temp.Submitted)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Insert func
func (repo *ReviewAnswerRepository) Insert(answer model.ReviewAnswer) (int, error) {
	var id int64

	query := "INSERT INTO user_reviews (user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, answer.UserReviewer, answer.UserTarget, answer.ReviewID, answer.AssignmentID,
		answer.Type, answer.Name, answer.Label, answer.Answer, answer.Comment, util.GetTimeInCorrectTimeZone())
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}

