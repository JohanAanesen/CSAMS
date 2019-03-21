package repositroy

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	_ "github.com/go-sql-driver/mysql" //database driver
)

// FormRepository struct
type FormRepository struct {
	db *sql.DB
}

// NewFormRepository func
func NewFormRepository(db *sql.DB) *FormRepository {
	return &FormRepository{
		db: db,
	}
}

// Fetch func
func (repo *FormRepository) Fetch(id int) (*model.Form, error) {
	result := model.Form{}
	query := "SELECT id, prefix, name, created FROM forms WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Prefix, &result.Name, &result.Created)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// FetchAll func
func (repo *FormRepository) FetchAll() ([]*model.Form, error) {
	result := make([]*model.Form, 0)
	query := "SELECT id, prefix, name, created FROM forms"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.Form

		err = rows.Scan(&temp.ID, &temp.Prefix, &temp.Name, &temp.Created)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Insert func
func (repo *FormRepository) Insert(form *model.Form) (int, error) {
	var id int64

	query := "INSERT INTO forms (prefix, name, created) VALUES (?, ?, ?)"
	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	form.Created = util.GetTimeInCorrectTimeZone()

	rows, err := tx.Exec(query, form.Prefix, form.Name, form.Created)
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

	fieldRepo := NewFieldRepository(repo.db)
	for _, item := range form.Fields {
		item.FormID = int(id)
		_, err := fieldRepo.Insert(&item)
		if err != nil {
			return 0, err
		}
	}

	return int(id), err
}

// Update func
func (repo *FormRepository) Update(id int, form *model.Form) error {
	if id != form.ID {
		return errors.New("form repository update: id does not match")
	}

	query := "UPDATE forms SET prefix = ?, name = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, form.Prefix, form.Name, form.ID)
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
func (repo *FormRepository) Delete(id int) error {
	query := "DELETE FROM forms WHERE id = ?"

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
