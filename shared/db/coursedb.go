package db

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
)

//GetCoursesToUser returns all the courses to the user
func GetCoursesToUser(userID int) model.Courses {

	//Create an empty courses array
	var courses model.Courses

	rows, err := GetDB().Query("SELECT course.* FROM course INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ?", userID)
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

		rows.Scan(&id, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Add course to courses array
		courses.Items = append(courses.Items, model.Course{
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
