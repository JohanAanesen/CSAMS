package repositroy

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	_ "github.com/go-sql-driver/mysql" //database driver
	"strings"
)

// FieldRepository struct
type FieldRepository struct {
	db *sql.DB
}

// NewFieldRepository func
func NewFieldRepository(db *sql.DB) *FieldRepository {
	return &FieldRepository{
		db: db,
	}
}

// Fetch func
func (repo *FieldRepository) Fetch(id int) (*model.Field, error) {
	result := model.Field{}
	query := "SELECT id, form_id, type, name, label, description, hasComment, priority, weight, choices FROM fields WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		var hasComment int
		var choices string

		err = rows.Scan(&result.ID, &result.FormID, &result.Type, &result.Name,
			&result.Label, &result.Description, &hasComment,
			&result.Order, &result.Weight, &choices)
		if err != nil {
			return &result, err
		}

		result.HasComment = hasComment == 1
		result.Choices = strings.Split(choices, ",")
	}

	return &result, err
}

// FetchAll func
func (repo *FieldRepository) FetchAll() ([]*model.Field, error) {
	result := make([]*model.Field, 0)
	query := "SELECT id, form_id, type, name, label, description, hasComment, priority, weight, choices FROM fields ORDER BY priority"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp = model.Field{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.FormID, &temp.Type, &temp.Name,
			&temp.Label, &temp.Description, &hasComment,
			&temp.Order, &temp.Weight, &choices)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, ",")

		result = append(result, &temp)
	}

	return result, err
}

// FetchAllFromForm func
func (repo *FieldRepository) FetchAllFromForm(formID int) ([]*model.Field, error) {
	result := make([]*model.Field, 0)
	query := "SELECT id, form_id, type, name, label, description, hasComment, priority, weight, choices FROM fields WHERE form_id = ? ORDER BY priority"

	rows, err := repo.db.Query(query, formID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp = model.Field{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.FormID, &temp.Type, &temp.Name,
			&temp.Label, &temp.Description, &hasComment,
			&temp.Order, &temp.Weight, &choices)
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
func (repo *FieldRepository) Insert(field *model.Field) (int, error) {
	var id int64

	query := "INSERT INTO fields (form_id, type, name, label, description, hasComment, priority, weight, choices) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	var hasComment = 0
	if field.HasComment {
		hasComment = 1
	}

	var choices = strings.Join(field.Choices, ",")

	rows, err := tx.Exec(query,
		field.FormID, field.Type, field.Name,
		field.Label, field.Description, hasComment,
		field.Order, field.Weight, choices)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}

// Update func
func (repo *FieldRepository) Update(id int, field *model.Field) error {
	if id != field.ID {
		return errors.New("field repository update: id does not match")
	}

	query := "UPDATE fields SET type = ?, name = ?, label = ?, description = ?, hasComment = ?, priority = ?, weight = ?, choices = ? WHERE id =?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	var hasComment = 0
	if field.HasComment {
		hasComment = 1
	}

	var choices = strings.Join(field.Choices, ",")

	_, err = tx.Exec(query, field.Type, field.Name, field.Label, field.Description, hasComment, field.Order, field.Weight, choices)
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

// DeleteAll func
func (repo *FieldRepository) DeleteAll(formID int) error {
	query := "DELETE FROM fields WHERE form_id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, formID)
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
