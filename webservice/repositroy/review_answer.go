package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
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

func fetchMany(repo *ReviewAnswerRepository, query string, args ...interface{}) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string
		var isRequired int

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight, &isRequired)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, "|")
		temp.Required = isRequired == 1

		result = append(result, &temp)
	}

	return result, err
}

// FetchForAssignment func
func (repo *ReviewAnswerRepository) FetchForAssignment(assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight, f.required FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.assignment_id = ?"
	return fetchMany(repo, query, assignmentID)
}

// FetchForTarget func
func (repo *ReviewAnswerRepository) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight, f.required FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_target = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, target, assignmentID)
	//return repo.FetchMany(query, target, assignmentID)
}

// FetchForReviewer func
func (repo *ReviewAnswerRepository) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight, f.required FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, reviewer, assignmentID)
}

// FetchForReviewerAndTarget func
func (repo *ReviewAnswerRepository) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight, f.required FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.user_target = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, reviewer, target, assignmentID)
}

// Insert func
func (repo *ReviewAnswerRepository) Insert(answer model.ReviewAnswer) (int, error) {
	var id int64

	query := "INSERT INTO user_reviews (user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	created := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	rows, err := tx.Exec(query, answer.UserReviewer, answer.UserTarget, answer.ReviewID, answer.AssignmentID,
		answer.Type, answer.Name, answer.Label, answer.Answer, answer.Comment, created)
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

// DeleteTarget func
func (repo *ReviewAnswerRepository) DeleteTarget(assignmentID, userID int) error {
	query := "DELETE FROM user_reviews WHERE assignment_id = ? AND user_target = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, assignmentID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// CountReviewsDone func
func (repo *ReviewAnswerRepository) CountReviewsDone(userID, assignmentID int) (int, error) {
	var result int

	query := "SELECT COUNT(DISTINCT user_target) FROM user_reviews WHERE user_reviewer = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, userID, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}

	return result, err
}

// MaxScore func
func (repo *ReviewAnswerRepository) MaxScore(assignmentID int) (int, error) {
	var result int

	query := `SELECT SUM(f.weight) FROM fields AS f INNER JOIN reviews AS r ON r.form_id = f.form_id INNER JOIN assignments AS a ON a.review_id = r.id WHERE a.id = ?`

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return 0, err
		}
	}

	return result, nil
}

// FetchRawReviewForUser func
func (repo *ReviewAnswerRepository) FetchRawReviewForUser(userID, assignmentID int) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	query := `SELECT ur.type, ur.answer, f.choices, f.weight FROM user_reviews AS ur
INNER JOIN fields AS f ON f.name = ur.name
INNER JOIN reviews AS r ON f.form_id = r.form_id
INNER JOIN assignments AS a ON a.review_id = r.id
WHERE a.id = ? AND ur.user_target = ? AND f.weight != 0
ORDER BY f.priority`

	rows, err := repo.db.Query(query, assignmentID, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var choices string

		err = rows.Scan(&temp.Type, &temp.Answer, &choices, &temp.Weight)
		if err != nil {
			return nil, err
		}

		temp.Choices = strings.Split(choices, "|")

		result = append(result, &temp)
	}

	return result, nil
}

// Update review answer and comment
func (repo *ReviewAnswerRepository) Update(targetID, reviewerID, assignmentID int, answer model.ReviewAnswer) error {
	query := "UPDATE user_reviews SET answer = ?, comment = ? WHERE user_reviewer = ? AND user_target = ? AND assignment_id = ? AND name = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, answer.Answer, answer.Comment, reviewerID, targetID, assignmentID, answer.Name)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
