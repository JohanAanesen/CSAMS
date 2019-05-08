package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// FAQRepository struct
type FAQRepository struct {
	db *sql.DB
}

// NewFAQRepository func
func NewFAQRepository(db *sql.DB) *FAQRepository {
	return &FAQRepository{
		db: db,
	}
}

// InsertNew inserts a new faq to db
func (repo *FAQRepository) InsertNew() error {

	questions := "# This FAQ is empty, press edit to fill in frequently asked questions"

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	query := "INSERT INTO `adminfaq` (`timestamp`, `questions`) VALUES (?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Exec(query, date, questions)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// Fetch fetches the faq in db
func (repo *FAQRepository) Fetch() (*model.Faq, error) {
	result := model.Faq{}

	query := "SELECT id, timestamp, questions FROM adminfaq"

	rows, err := repo.db.Query(query)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Timestamp, &result.Questions)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// Update updates the faq in db
func (repo *FAQRepository) Update(questions string) error {

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	query := "UPDATE `adminfaq` SET `timestamp` = ?, `questions` = ?"

	//insert into database
	rows, err := repo.db.Query(query, date, questions)

	// Check for errors
	if err != nil {
		return err
	}

	defer rows.Close()

	return err
}
