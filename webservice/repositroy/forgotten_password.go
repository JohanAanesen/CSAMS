package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
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
	var id int

	query := "INSERT INTO `forgotten_password` (`hash`, `user_id`, `timestamp`) VALUES (?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return id, err
	}

	_, err = tx.Exec(query, forgottenPass.Hash, forgottenPass.UserID, forgottenPass.TimeStamp)
	if err != nil {
		tx.Rollback()
		return id, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return id, err
	}

	return id, nil
}

// Match checks if the hash exists in the db
func (repo *ForgottenPassRepository) Match(hash string) (bool, error) {

	query := "SELECT COUNT(id) FROM `forgotten_password` WHERE `hash` = ?"

	rows, err := repo.db.Query(query, hash)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var count int

		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}

		// If count is over 0, the hash exists in the db
		if count > 0 {
			return true, nil
		}
	}

	return false, err
}
