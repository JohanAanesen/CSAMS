package repository

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	_ "github.com/go-sql-driver/mysql" //database driver
)

// ReviewRepository struct
type ReviewRepository struct {
	db *sql.DB
}

// NewReviewRepository func
func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

// FetchAll func
func (repo *ReviewRepository) FetchAll() ([]*model.Review, error) {
	result := make([]*model.Review, 0)
	query := "SELECT id, form_id FROM reviews"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.Review

		err = rows.Scan(&temp.ID, &temp.FormID)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Fetch func
func (repo *ReviewRepository) Fetch(id int) (*model.Review, error) {
	result := model.Review{}
	query := "SELECT id, form_id FROM reviews WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.FormID)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// Insert func
func (repo *ReviewRepository) Insert(form model.Form) (int, error) {
	formRepo := NewFormRepository(repo.db)
	formID, err := formRepo.Insert(&form)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO reviews (form_id) VALUES (?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	rows, err := tx.Exec(query, formID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := rows.LastInsertId()
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

// Update func
func (repo *ReviewRepository) Update(id int, review model.Review) error {
	if id != review.ID {
		return errors.New("review repository update: id does not match")
	}

	query := "UPDATE forms SET form_id = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, review.FormID)
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

// DeleteByFormID func
func (repo *ReviewRepository) DeleteByFormID(id int) error {
	query := "DELETE FROM reviews WHERE form_id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id)
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

// FetchReviewUsers func
func (repo *ReviewRepository) FetchReviewUsers(userID, assignmentID int) ([]*model.User, error) {
	result := make([]*model.User, 0)
	query := "SELECT pr.review_user_id, u.name FROM peer_reviews AS pr INNER JOIN users AS u ON pr.review_user_id = u.id WHERE pr.user_id = ? AND pr.assignment_id = ?"

	rows, err := repo.db.Query(query, userID, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var temp model.User

		err = rows.Scan(&temp.ID, &temp.Name)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// IsUserTheReviewer func
func (repo *ReviewRepository) IsUserTheReviewer(reviewer int, target int, assignment int) (bool, error) {
	result := make([]int, 0)
	query := "SELECT id FROM peer_reviews WHERE user_id = ? AND review_user_id = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, reviewer, target, assignment)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp int

		err = rows.Scan(&temp)
		if err != nil {
			return false, err
		}

		result = append(result, temp)
	}

	return len(result) > 0, err
}

// IsUsed func
func (repo *ReviewRepository) IsUsed(id int) (bool, error) {
	query := "SELECT r.form_id FROM assignments AS a INNER JOIN reviews AS r ON a.review_id = r.id"

	rows, err := repo.db.Query(query)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var temp int
		err = rows.Scan(&temp)
		if err != nil {
			return false, err
		}

		if temp == id {
			return true, nil
		}
	}

	return false, err
}
