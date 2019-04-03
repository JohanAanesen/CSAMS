package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// ValidationRepository struct
type ValidationRepository struct {
	db *sql.DB
}

// NewValidationRepository func
func NewValidationRepository(db *sql.DB) *ValidationRepository {
	return &ValidationRepository{
		db: db,
	}
}

// Insert inserts a new forgottenpass in the db
func (repo *ValidationRepository) Insert(validation model.Validation) (int, error) {
	var id int64

	query := "INSERT INTO `validation` (`hash`, `user_id`, `timestamp`) VALUES (?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, validation.Hash, validation.UserID, util.ConvertTimeStampToString(validation.TimeStamp))
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

	return int(id), nil
}

// FetchAll fetches all validation rows
func (repo *ValidationRepository) FetchAll() ([]*model.Validation, error) {
	result := make([]*model.Validation, 0)

	query := "SELECT `id`, `hash`, `user_id`, `valid`, `timestamp`  FROM `validation`"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Validation{}

		err = rows.Scan(&temp.ID, &temp.Hash, &temp.UserID, &temp.Valid, &temp.TimeStamp)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// UpdateValidation updates the validation column
func (repo *ValidationRepository) UpdateValidation(id int, state bool) error {
	query := "UPDATE `validation` SET `valid` = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, state, id)
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
