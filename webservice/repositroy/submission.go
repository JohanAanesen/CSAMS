package repositroy

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	_ "github.com/go-sql-driver/mysql" //database driver
)

// SubmissionRepository struct
type SubmissionRepository struct {
	db *sql.DB
}

// NewSubmissionRepository func
func NewSubmissionRepository(db *sql.DB) *SubmissionRepository {
	return &SubmissionRepository{
		db: db,
	}
}

// FetchAll func
func (repo *SubmissionRepository) FetchAll() ([]*model.Submission, error) {
	result := make([]*model.Submission, 0)
	query := "SELECT id, form_id FROM submissions"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.Submission

		err = rows.Scan(&temp.ID, &temp.FormID)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Fetch func
func (repo *SubmissionRepository) Fetch(id int) (*model.Submission, error) {
	result := model.Submission{}
	query := "SELECT id, form_id FROM submissions WHERE id = ?"

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
func (repo *SubmissionRepository) Insert(form model.Form) (int, error) {
	formRepo := NewFormRepository(repo.db)
	formID, err := formRepo.Insert(&form)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO submissions (form_id) VALUES (?)"

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
func (repo *SubmissionRepository) Update(id int, submission model.Submission) error {
	if id != submission.ID {
		return errors.New("review repository update: id does not match")
	}

	query := "UPDATE submissions SET form_id = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, submission.FormID)
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

// Delete func
func (repo *SubmissionRepository) Delete(id int) error {
	query := "DELETE FROM submissions WHERE id = ?"

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
	}

	return err
}

// InUse func
func (repo *SubmissionRepository) IsUsed(id int) (bool, error) {
	query := "SELECT s.form_id FROM assignments AS a INNER JOIN submissions AS s ON a.review_id = s.id"

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