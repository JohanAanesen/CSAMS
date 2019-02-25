package model

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"log"
)

// Courses hold the data for a slice of Course-struct
type Courses struct {
	Items []Course `json:"courses"`
}

// Course holds the data for courses
type Course struct {
	ID          int          `json:"id"`
	Hash        string       `json:"hash"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Teacher     int          `json:"teacher"`
	Year        string       `json:"year"`
	Semester    string       `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}

//GetCoursesToUser returns all the courses to the user
func GetCoursesToUser(userID int) (Courses, error) {

	// Create an empty courses array
	var courses Courses

	// Query gets all courses to user with userid: userID and sort them by year and semester
	rows, err := db.GetDB().Query("SELECT course.* FROM course INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ? ORDER BY course.year DESC, course.semester DESC", userID)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course array if it fails
		return courses, err
	}

	for rows.Next() {
		var id int
		var hash string
		var courseCode string
		var courseName string
		var teacher int
		var description string
		var year string
		var semester string

		rows.Scan(&id, &hash, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Add course to courses array
		courses.Items = append(courses.Items, Course{
			ID:          id,
			Hash:        hash,
			Code:        courseCode,
			Name:        courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		})
	}

	return courses, nil
}

//GetCourse returns a given course
func GetCourse(courseID int) Course {
	//Create an empty courses array
	var course Course

	rows, err := db.GetDB().Query("SELECT course.* FROM course WHERE course.id = ?", courseID)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course array if it fails
		return Course{}
	}

	for rows.Next() {
		var id int
		var hash string
		var courseCode string
		var courseName string
		var teacher int
		var description string
		var year string
		var semester string

		rows.Scan(&id, &hash, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Add course to courses array
		course = Course{
			ID:          id,
			Hash:        hash,
			Code:        courseCode,
			Name:        courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		}
	}

	return course
}

// CourseExists checks if the course exists in the database
func CourseExists(hash string) Course {
	rows, err := db.GetDB().Query("SELECT course.* FROM course WHERE hash = ?", hash)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course if it fails
		return Course{ID: -1}
	}

	for rows.Next() {
		var id int
		var hash string
		var courseCode string
		var courseName string
		var teacher int
		var description string
		var year string
		var semester string

		rows.Scan(&id, &hash, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Fill course object/struct
		return Course{
			ID:          id,
			Hash:        hash,
			Code:        courseCode,
			Name:        courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		}
	}

	return Course{ID: -1}
}

// UserExistsInCourse checks if user exists in course
func UserExistsInCourse(userID int, courseID int) bool {

	// Checks if user exists in course
	rows, err := db.GetDB().Query("SELECT * FROM usercourse WHERE userid = ? AND courseid = ?", userID, courseID)

	// return false if not
	if err != nil {
		log.Println(err.Error())

		return false
	}

	for rows.Next() {
		var userid int
		var courseid string

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
	rows, err := db.GetDB().Query("INSERT INTO `usercourse` (`userid`, `courseid`) VALUES (?, ?)", userID, courseID)

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
