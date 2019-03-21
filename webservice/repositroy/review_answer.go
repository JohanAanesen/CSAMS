package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"strings"
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

	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, ",")

		result = append(result, &temp)
	}

	return result, err
}

// FetchForTarget func
func (repo *ReviewAnswerRepository) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_target = ? AND ur.assignment_id = ?"

	rows, err := repo.db.Query(query, target, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, ",")

		result = append(result, &temp)
	}

	return result, err
}

// FetchForReviewer func
func (repo *ReviewAnswerRepository) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.assignment_id = ?"

	rows, err := repo.db.Query(query, reviewer, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, ",")

		result = append(result, &temp)
	}

	return result, err
}

// FetchForReviewerAndTarget func
func (repo *ReviewAnswerRepository) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.user_target = ? AND ur.assignment_id = ?"

	rows, err := repo.db.Query(query, reviewer, target, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, ",")

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
