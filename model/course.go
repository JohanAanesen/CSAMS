package model

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
)

// Courses hold the data for a slice of Course-struct
type Courses struct {
	Items []Course `json:"courses"`
}

// Course holds the data for courses
type Course struct {
	ID          int          `json:"id"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Teacher     int          `json:"teacher"`
	Year        string       `json:"year"`
	Semester    string       `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}


// TODO (Svein): make this a method of a struct (From: coursedb.go)
//GetCoursesToUser returns all the courses to the user
func GetCoursesToUser(userID int) Courses {

	//Create an empty courses array
	var courses Courses

	rows, err := db.GetDB().Query("SELECT course.* FROM course INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ?", userID)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course array if it fails
		return courses
	}

	for rows.Next() {
		var id int
		var courseCode string
		var courseName string
		var teacher int
		var description string
		var year string
		var semester string

		// TODO (Svein): error-handling
		rows.Scan(&id, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Add course to courses array
		courses.Items = append(courses.Items, Course{
			ID:          id,
			Code:        courseCode,
			Name:        courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		})
	}

	return courses
}