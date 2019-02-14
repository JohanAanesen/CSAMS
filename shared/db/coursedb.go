package db

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"log"
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

// CourseExists checks if the course exists in the database
func CourseExists(uniqueID string) model.Course {
	rows, err := GetDB().Query("SELECT course.* FROM course WHERE id = ?", uniqueID)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course if it fails
		return model.Course{ID: -1}
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

		// Fill course object/struct
		return model.Course{
			ID:          id,
			Code:        courseCode,
			Name:        courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		}
	}

	return model.Course{ID: -1}
}

// UserExistsInCourse checks if user exists in course
func UserExistsInCourse(userID int, courseID int) bool {

	// Checks if user exists in course
	rows, err := GetDB().Query("SELECT * FROM usercourse WHERE userid = ? AND courseid = ?", userID, courseID)

	// return false if not
	if err != nil {
		log.Println(err.Error())

		return false
	}

	for rows.Next() {
		var userid int
		var courseid int

		rows.Scan(&userid, &courseid)

		// return true if user exists in course
		return true
	}

	// default: return false
	return false
}

// AddUserToCourse adds the user to a course
func AddUserToCourse(userID int, courseID int) bool {

	// Sql query
	rows, err := GetDB().Query("INSERT INTO `usercourse` (`userid`, `courseid`) VALUES (?, ?)", userID, courseID)

	// Return false if error
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	// Close this
	defer rows.Close()

	// return true
	return true
}
