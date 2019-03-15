package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"log"
	"time"
)

// Faq Struct for keeping the frequently asked questions under /admin/faq
type Faq struct {
	Date      time.Time // Last edited time
	Questions string    // The markdown with questions and answers
}

// GetDateAndQuestionsFAQ returns the date and question from the faq
func GetDateAndQuestionsFAQ() Faq {
	content := Faq{Questions: "-1"}

	//insert into database
	rows, err := db.GetDB().Query("SELECT timestamp, questions FROM `adminfaq` WHERE id = 1") // OBS! ID is always 1 since it's only one entry in the table

	// Log error
	if err != nil {
		log.Println(err.Error())
		return content
	}

	for rows.Next() {
		var timestamp time.Time
		var questions string

		err = rows.Scan(&timestamp, &questions)
		if err != nil {
			return Faq{}
		}

		content = Faq{
			Date:      timestamp,
			Questions: questions,
		}
	}

	return content

}

// UpdateFAQ updates the questions and date in FAQ
func UpdateFAQ(newFaq string) error {

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	// Update to database
	rows, err := db.GetDB().Query("UPDATE `adminfaq` SET `timestamp` = ?, `questions` = ? WHERE `id` = 1", date, newFaq)

	// Log error if it exists
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil

}
