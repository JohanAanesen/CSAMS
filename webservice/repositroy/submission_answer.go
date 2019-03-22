package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"strings"
)

// SubmissionAnswerRepository struct
type SubmissionAnswerRepository struct {
	db *sql.DB
}

// NewSubmissionAnswerRepository func
func NewSubmissionAnswerRepository(db *sql.DB) *SubmissionAnswerRepository {
	return &SubmissionAnswerRepository{
		db: db,
	}
}

// Fetch func
func (repo *SubmissionAnswerRepository) Fetch(id int) (*model.SubmissionAnswer, error) {
	result := model.SubmissionAnswer{}

	query := "SELECT id, user_id, assignment_id, submission_id, type, name, label, answer, comment, submitted FROM user_submissions WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.UserID, &result.AssignmentID, &result.SubmissionID,
			&result.Type, &result.Name, &result.Label, &result.Answer, &result.Comment, &result.Submitted)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// FetchAll func
func (repo *SubmissionAnswerRepository) FetchAll() ([]*model.SubmissionAnswer, error) {
	result := make([]*model.SubmissionAnswer, 0)
	query := "SELECT id, user_id, assignment_id, submission_id, type, name, label, answer, comment, submitted FROM user_submissions"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.SubmissionAnswer{}

		err = rows.Scan(&temp.ID, &temp.UserID, &temp.AssignmentID, &temp.SubmissionID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Answer, &temp.Comment, &temp.Submitted)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchAllForUserAndAssignment func
func (repo *SubmissionAnswerRepository) FetchAllForUserAndAssignment(userID, assignmentID int) ([]*model.SubmissionAnswer, error) {
	result := make([]*model.SubmissionAnswer, 0)

	query := "SELECT us.id, us.user_id, us.assignment_id, us.submission_id, f.type, f.name, f.label, f.description, us.answer, us.comment, us.submitted, f.hasComment, f.choices, f.weight FROM user_submissions AS us INNER JOIN fields AS f ON us.name = f.name WHERE us.user_id = ? AND us.assignment_id = ?"

	rows, err := repo.db.Query(query, userID, assignmentID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.SubmissionAnswer{}
		var choices string
		var hasComment int

		err = rows.Scan(&temp.ID, &temp.UserID, &temp.AssignmentID, &temp.SubmissionID,
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
func (repo *SubmissionAnswerRepository) Insert(answer model.SubmissionAnswer) (int, error) {
	var id int64

	query := "INSERT INTO user_submissions (user_id, assignment_id, submission_id, type, name, label, answer, comment, submitted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	submitted := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())
	rows, err := tx.Exec(query, answer.UserID, answer.AssignmentID,
		answer.SubmissionID, answer.Type, answer.Name, answer.Label, answer.Answer, answer.Comment, submitted)
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

// CountForAssignment func
func (repo *SubmissionAnswerRepository) CountForAssignment(assignmentID int) (int, error) {
	var count int

	query := "SELECT COUNT(DISTINCT user_id) FROM user_submissions WHERE assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return count, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, err
}

// Update func
func (repo *SubmissionAnswerRepository) Update(answer model.SubmissionAnswer) error {
	query := "UPDATE user_submissions SET answer = ?, comment = ?, submitted = ? WHERE name = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	submitted := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())
	_, err = tx.Exec(query, answer.Answer, answer.Comment, submitted, answer.Name)
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

// DeleteFromAssignment func
func (repo *SubmissionAnswerRepository) DeleteFromAssignment(assignmentID int) error {
	query := "DELETE FROM user_submissions WHERE assignment_id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, assignmentID)
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
