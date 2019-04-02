package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
)

// ForgottenPassRepository struct
type ForgottenPassRepository struct {
	db *sql.DB
}

// NewForgottenPassRepository func
func NewForgottenPassRepository(db *sql.DB) *ForgottenPassRepository {
	return &ForgottenPassRepository{
		db: db,
	}
}

// Insert inserts a new forgottenpass in the db
func (repo *ForgottenPassRepository) Insert(forgottenPass model.ForgottenPass) (int, error) {
	var id int64

	query := "INSERT INTO `forgotten_password` (`hash`, `user_id`, `timestamp`) VALUES (?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, forgottenPass.Hash, forgottenPass.UserID, util.ConvertTimeStampToString(forgottenPass.TimeStamp))
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

// FetchAll fetches all forgotten_password rows
func (repo *ForgottenPassRepository) FetchAll() ([]*model.ForgottenPass, error) {
	result := make([]*model.ForgottenPass, 0)

	query := "SELECT `id`, `hash`, `user_id`, `timestamp`  FROM `forgotten_password`"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ForgottenPass{}

		err = rows.Scan(&temp.ID, &temp.Hash, &temp.UserID, &temp.TimeStamp)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}
