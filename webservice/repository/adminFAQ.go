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

// Fetch fetches the faq in db
func (repo *FAQRepository) Fetch() (*model.Faq, error) {
	id := 1
	result := model.Faq{}

	query := "SELECT id, timestamp, questions FROM adminfaq WHERE id = ?"

	rows, err := repo.db.Query(query, id)
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

	query := "UPDATE `adminfaq` SET `timestamp` = ?, `questions` = ? WHERE `id` = ?"

	//insert into database
	rows, err := repo.db.Query(query, date, questions, 1) // OBS! ID is always 1 since it's only one entry in the table

	// Check for errors
	if err != nil {
		return err
	}

	defer rows.Close()

	return err
}
