package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"log"
	"time"
)

// Struct for keeping the frequently asked questions under /admin/faq
type Faq struct {
	Date      time.Time // Last edited time
	Questions string    // The markdown with questions and answers
}

func GetDateAndQuestionsFAQ() Faq {
	content := Faq{Questions: "-1"}

	// TODO : it feels wrong to have this here, but I think this is correct :S

	//insert into database
	rows, err := db.GetDB().Query("SELECT timestamp, questions FROM `adminfaq` WHERE id = 1") // OBS! Id is always 1 since it's only one entry in the table

	// Log error
	if err != nil {
		log.Println(err.Error())
		return content
	}

	for rows.Next() {
		var timestamp time.Time
		var questions string

		rows.Scan(&timestamp, &questions)

		content = Faq{
			Date:      timestamp,
			Questions: questions,
		}
	}

	return content

}
