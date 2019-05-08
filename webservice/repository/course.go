package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

// CourseRepository represents the layer between the database and the Course-model
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository returns a pointer to a new CourseRepository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Fetch a single course from the database
func (repo *CourseRepository) Fetch(id int) (*model.Course, error) {
	// Initialize an empty course
	result := model.Course{}
	// Query string
	query := "SELECT id, hash, coursecode, coursename, teacher, description, year, semester FROM course WHERE id = ?"
	// Perform query-string with id as a parameter
	rows, err := repo.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Scan all columns from row
		err = rows.Scan(&result.ID, &result.Hash, &result.Code,
			&result.Name, &result.Teacher, &result.Description,
			&result.Year, &result.Semester)
		if err != nil {
			return nil, err
		}
	}
	// Return result
	return &result, nil
}

// FetchAll courses from the database
func (repo *CourseRepository) FetchAll() ([]*model.Course, error) {
	// Create empty course slice
	result := make([]*model.Course, 0)
	// Query string
	query := "SELECT id, hash, coursecode, coursename, teacher, description, year, semester FROM course"
	// Perform query
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Temporary course object
		temp := model.Course{}
		// Scan columns from row
		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return nil, err
		}
		// Append temporary course to result
		result = append(result, &temp)
	}
	// Return result
	return result, nil
}

// FetchAllForUserOrdered all courses for a user, ordered
func (repo *CourseRepository) FetchAllForUserOrdered(userID int) ([]*model.Course, error) {
	// Create empty course slice
	result := make([]*model.Course, 0)
	// Query string
	query := "SELECT c.id, c.hash, c.coursecode, c.coursename, c.teacher, c.description, c.year, c.semester FROM course AS c INNER JOIN usercourse AS uc ON c.id = uc.courseid WHERE uc.userid = ? ORDER BY c.year DESC, c.semester ASC, c.coursename ASC"
	// Perform query
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Create temporary course object
		temp := model.Course{}
		// Scan columns from the row
		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return nil, err
		}
		// Append temporary object to result
		result = append(result, &temp)
	}
	// Return result
	return result, nil
}

// FetchAllStudentsFromCourse func
func (repo *CourseRepository) FetchAllStudentsFromCourse(courseID int) ([]*model.User, error) {
	result := make([]*model.User, 0)

	query := "SELECT u.id, u.name, u.email_student, u.teacher, u.email_private FROM users AS u INNER JOIN usercourse AS uc ON u.id = uc.userid WHERE uc.courseid = ? AND u.teacher = 0"

	rows, err := repo.db.Query(query, courseID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.User{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.EmailStudent,
			&temp.Teacher, &temp.EmailPrivate)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Insert course to the database
func (repo *CourseRepository) Insert(course model.Course) (int, error) {
	// Integer to hold the id of last inserted row
	var id int64
	// Query string
	query := "INSERT INTO course (hash, coursecode, coursename, teacher, description, year, semester) VALUES (?, ?, ?, ?, ?, ?, ?)"
	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}
	// Execute query with parameters
	rows, err := tx.Exec(query, course.Hash, course.Code,
		course.Name, course.Teacher, course.Description,
		course.Year, course.Semester)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}
	// Get last inserted ID
	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}
	// Return result
	return int(id), nil
}

// InsertUser to a course, gives a relationship between a user and a course in the database
func (repo *CourseRepository) InsertUser(userID int, courseID int) error {
	// Query string
	query := "INSERT INTO usercourse (userid, courseid) VALUES (?, ?)"
	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	// Execute query with parameters
	_, err = tx.Exec(query, userID, courseID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	// Return no error
	return nil
}

// UserInCourse checks if user is in given course
func (repo *CourseRepository) UserInCourse(userID, courseID int) (bool, error) {
	// Query string
	query := "SELECT COUNT(DISTINCT userid) FROM usercourse WHERE userid =? AND courseid = ?"
	// Perform query with parameters
	rows, err := repo.db.Query(query, userID, courseID)
	if err != nil {
		return false, err
	}
	// Close connection
	defer rows.Close()
	// Check if it returned any rows
	if rows.Next() {
		// Integer for storing the result
		var result int
		// Scan column from the row
		err = rows.Scan(&result)
		if err != nil {
			return false, err
		}

		// If the user is in the course
		if result == 1 {
			return true, nil
		}
	}

	// Return false
	return false, nil
}

// Update func
func (repo *CourseRepository) Update(course model.Course) error {
	query := "UPDATE course SET coursecode = ?, coursename = ?, description = ?, semester = ? WHERE id = ?"
	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, course.Code, course.Name, course.Description, course.Semester, course.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	// All good, return no error
	return nil
}

// Delete course from the database, based on course ID
func (repo *CourseRepository) Delete(id int) error {
	// Query string
	query := "DELETE FROM course WHERE id = ?"
	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	// Execute query with parameter
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	// All good, no error
	return nil
}

// RemoveUser removes the relationship between a user and a course in the database
func (repo *CourseRepository) RemoveUser(userID, courseID int) error {
	// Query string
	query := "DELETE FROM usercourse WHERE userid = ? AND courseid = ?"
	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	// Execute query with parameters
	_, err = tx.Exec(query, userID, courseID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	// All good, no error
	return nil
}

// FetchAllForUser fetches all courses for a given user
func (repo *CourseRepository) FetchAllForUser(userID int) ([]*model.Course, error) {
	// Empty course slice
	result := make([]*model.Course, 0)
	// Query string
	query := "SELECT c.id, c.hash, c.coursecode, c.coursename, c.teacher, c.description, c.year, c.semester FROM course AS c INNER JOIN usercourse AS uc ON c.id = uc.courseid WHERE uc.userid = ? ORDER BY c.year DESC, c.semester ASC, c.coursename ASC"
	// Perform query
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Temporary course object
		temp := model.Course{}
		// Scan columns from row
		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return nil, err
		}
		// Append temporary object to result
		result = append(result, &temp)
	}
	// Return result
	return result, nil
}
