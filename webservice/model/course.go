package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"log"
)

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

// CourseRepository holds all assignments, and DB-functions
type CourseRepository struct{}

// GetSingle retrieves a single Course based on Primary Key (id)
func (repo *CourseRepository) GetSingle(id int) (Course, error) {
	// Declare empty struct
	var result Course

	// Create query string
	query := "SELECT id, hash, coursecode, coursename, description, teacher, year, semester FROM course WHERE id = ?"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, id)
	// Check for error
	if err != nil {
		return Course{}, err
	}
	// Close connection
	defer rows.Close()

	for rows.Next() {
		// Scan row for data
		err = rows.Scan(&result.ID, &result.Hash, &result.Code,
			&result.Name, &result.Description, &result.Teacher, &result.Year,
			&result.Semester)
		// Check for error
		if err != nil {
			return Course{}, err
		}
	}

	return result, nil
}

// GetAll returns all Courses in the database
func (repo *CourseRepository) GetAll() ([]Course, error) {
	// Declare empty slice
	var result []Course

	// Create query string
	query := "SELECT id, hash, coursecode, coursename, description, teacher, year, semester FROM course;"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query)
	if err != nil {
		return nil, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		// Declare empty struct
		var course Course
		// Scan rows
		err := rows.Scan(&course.ID, &course.Hash, &course.Code,
			&course.Name, &course.Description, &course.Teacher, &course.Year,
			&course.Semester)
		// Check for error
		if err != nil {
			return nil, err
		}

		// Append retrieved row
		result = append(result, course)
	}

	return result, nil
}

// GetAllToUserSorted Gets all courses to user and returns them sorted by year
func (repo *CourseRepository) GetAllToUserSorted(UserID int) ([]Course, error) {

	// Declare empty slice
	var result []Course

	// Create query string
	// The tables is connected like this example: users -> usercourse -> course
	query := "SELECT course.id, course.hash, course.coursecode, course.coursename, course.description, course.teacher, course.year, course.semester  " +
		"FROM `course` INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ? " +
		"ORDER BY course.year DESC, course.semester ASC, course.coursename ASC;"

	// Prepare and execute query
	rows, err := db.GetDB().Query(query, UserID)
	if err != nil {
		return nil, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		// Declare empty struct
		var course Course
		// Scan rows
		err := rows.Scan(&course.ID, &course.Hash, &course.Code,
			&course.Name, &course.Description, &course.Teacher, &course.Year,
			&course.Semester)
		// Check for error
		if err != nil {
			return nil, err
		}

		// Append retrieved row
		result = append(result, course)
	}

	return result, nil
}

// Update an course based on the ID and the data inside an Course-object
func (repo *CourseRepository) Update(id int, course Course) error {
	query := "UPDATE course SET coursecode=?, coursename=?, description=?, semester=? WHERE id=?"

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, course.Code, course.Name, course.Description, course.Semester, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

// RemoveUser removes a user from a course
func (repo *CourseRepository) RemoveUser(courseID int, userID int) error {

	// Create query string
	query := "DELETE from usercourse WHERE userid = ? AND courseid = ?"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, userID, courseID)

	// Check for error
	if err != nil {
		return err
	}
	// Close connection
	defer rows.Close()

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO Brede : Change ex `var courseCOde string` to only use `course.Code`

/*
// GetCourseCodeAndName gets the course code and name to the course
func GetCourseCodeAndName(courseID int) (Course, error) {
	//Create an empty courses array
	var course Course

	rows, err := db.GetDB().Query("SELECT course.coursecode, course.coursename FROM course WHERE course.id = ?", courseID)
	if err != nil {
		log.Println(err.Error())

		// returns empty course array if it fails
		return course, err
	}

	for rows.Next() {
		var courseCode string
		var courseName string

		err := rows.Scan(&courseCode, &courseName)
		if err != nil {
			return course, err
		}

		// Add course to courses array
		course = Course{
			Code: courseCode,
			Name: courseName,
		}
	}

	return course, nil
}

//GetCoursesToUser returns all the courses to the user
func GetCoursesToUser(userID int) ([]Course, error) {

	// Create an empty courses array
	var courses []Course

	// Query gets all courses to user with userid: userID and sort them by year and semester
	rows, err := db.GetDB().Query("SELECT course.* FROM course INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ? ORDER BY course.year DESC, course.semester DESC", userID)
	if err != nil {
		log.Println(err.Error())

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
		courses = append(courses, Course{
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
		log.Println(err.Error())

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

*/

// CourseExists checks if the course exists in the database
func (repo *CourseRepository) CourseExists(hash string) Course {
	rows, err := db.GetDB().Query("SELECT course.* FROM course WHERE hash = ?", hash)
	if err != nil {
		log.Println(err.Error())

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
func (repo *CourseRepository) UserExistsInCourse(userID int, courseID int) bool {

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
func (repo *CourseRepository) AddUserToCourse(userID int, courseID int) bool {

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
