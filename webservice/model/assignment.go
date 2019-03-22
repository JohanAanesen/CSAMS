package model

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	_ "github.com/go-sql-driver/mysql" //database driver
	"log"
	"time"
)

// Assignment hold the data for a single assignment
type Assignment struct {
	ID           int           `json:"id" db:"id"`
	Name         string        `json:"name" db:"name"`
	Description  string        `json:"description" db:"description"`
	Created      time.Time     `json:"created" db:"created"`
	Publish      time.Time     `json:"publish" db:"publish"`
	Deadline     time.Time     `json:"deadline" db:"deadline"`
	CourseID     int           `json:"course_id" db:"course_id"`
	SubmissionID sql.NullInt64 `json:"-" db:"submission_id"`
	ReviewID     sql.NullInt64 `json:"-" db:"review_id"`
	Submission   Submission    `json:"submission"`
	Review       Review        `json:"review"`
	Reviewers    sql.NullInt64 `json:"reviewers"`
	ValidationID sql.NullInt64 `json:"validation_id"`
}

// AssignmentRepository holds all assignments, and DB-functions
type AssignmentRepository struct{}

// GetSingle retrieves a single Assignment based on Primary Key (id)
func (repo *AssignmentRepository) GetSingle(id int) (Assignment, error) {
	// Declare empty struct
	var result Assignment

	// Select query string
	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, reviewers FROM assignments WHERE id = ?"
	// Prepare and execute query
	rows, err := db.GetDB().Query(query, id)
	// Check for error
	if err != nil {
		return Assignment{}, err
	}
	// Close connection
	defer rows.Close()

	for rows.Next() {
		// Scan row for data
		err = rows.Scan(&result.ID, &result.Name, &result.Description,
			&result.Created, &result.Publish, &result.Deadline, &result.CourseID,
			&result.SubmissionID, &result.ReviewID, &result.Reviewers)
		// Check for error
		if err != nil {
			return Assignment{}, err
		}
	}

	return result, nil
}

// GetAll returns all assignments in the database
func (repo *AssignmentRepository) GetAll() ([]Assignment, error) {
	// Declare empty slice
	var result []Assignment

	// Select query string
	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, reviewers FROM assignments;"
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
		var assignment Assignment
		// Scan rows
		err := rows.Scan(&assignment.ID, &assignment.Name, &assignment.Description,
			&assignment.Created, &assignment.Publish, &assignment.Deadline,
			&assignment.CourseID, &assignment.SubmissionID, &assignment.ReviewID)
		// Check for error
		if err != nil {
			return nil, err
		}

		// Append retrieved row
		result = append(result, assignment)
	}

	return result, nil
}

// GetAllToUserSorted Gets all assignment to user and returns them sorted by deadline
func (repo *AssignmentRepository) GetAllToUserSorted(UserID int) ([]Assignment, error) {

	// Declare empty slice
	var result []Assignment

	// Select query string
	// The tables is connected like this example: users -> usercourse -> course -> assignments
	query := "SELECT assignments.id, assignments.name, assignments.description, assignments.created, assignments.publish, assignments.deadline, assignments.course_id  " +
		"FROM `assignments` INNER JOIN course ON assignments.course_id = course.id " +
		"INNER JOIN usercourse ON usercourse.courseid = course.id WHERE usercourse.userid = ? " +
		"AND usercourse.courseid = assignments.course_id ORDER BY assignments.deadline;"

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
		var assignment Assignment
		// Scan rows
		err := rows.Scan(&assignment.ID, &assignment.Name, &assignment.Description,
			&assignment.Created, &assignment.Publish, &assignment.Deadline,
			&assignment.CourseID)
		// Check for error
		if err != nil {
			return nil, err
		}

		// Append retrieved row
		result = append(result, assignment)
	}

	return result, nil
}

// Insert a new assignment to the database
func (repo *AssignmentRepository) Insert(assignment Assignment) (int, error) {

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	// Insert query string
	query := "INSERT INTO assignments (name, description, publish, deadline, course_id) VALUES (?, ?, ?, ?, ?);"
	// Prepare and execute query
	ex, err := tx.Exec(query, assignment.Name, assignment.Description, assignment.Publish, assignment.Deadline, assignment.CourseID)
	// Check for error
	if err != nil {
		tx.Rollback() //quit transaction if error
		return 0, err
	}

	// Get last inserted ID
	id, err := ex.LastInsertId()
	// Check for error
	if err != nil {
		tx.Rollback() //quit transaction if error
		return 0, err
	}

	// Check if we have set a submission_id
	if assignment.SubmissionID.Valid {
		// Update query string
		query := "UPDATE assignments SET submission_id = ? WHERE id = ?;"
		// Prepare and execute query
		_, err := tx.Exec(query, assignment.SubmissionID, id)
		// Check for error
		if err != nil {
			tx.Rollback() //quit transaction if error
			return 0, err
		}
	}

	// Check if we have set a review_id
	if assignment.ReviewID.Valid {
		// Update query string
		query := "UPDATE assignments SET review_id = ? WHERE id = ?;"
		// Prepare and execute query
		_, err := tx.Exec(query, assignment.ReviewID, id)
		// Check for error
		if err != nil {
			tx.Rollback() //quit transaction if error
			return 0, err
		}
	}

	// Check if we have set reviewers
	if assignment.Reviewers.Valid {
		// Update query string
		query := "UPDATE assignments SET reviewers = ? WHERE id = ?;"
		// Prepare and execute query
		_, err := tx.Exec(query, assignment.Reviewers, id)
		// Check for error
		if err != nil {
			tx.Rollback() //quit transaction if error
			return 0, err
		}
	}

	// Set created date
	query = "UPDATE assignments SET created = ? WHERE id = ?;"
	// Prepare and execute query
	_, err = tx.Exec(query, date, id)
	// Check for error
	if err != nil {
		tx.Rollback() //quit transaction if error
		return 0, err
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Update an assignment based on the ID and the data inside an Assignment-object
func (repo *AssignmentRepository) Update(id int, assignment Assignment) error {
	query := "UPDATE assignments SET name=?, description=?, course_id=?, submission_id=?, publish=?, deadline=?, reviewers=? WHERE id=?"

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	if assignment.SubmissionID.Valid {
		_, err = tx.Exec(query, assignment.Name, assignment.Description, assignment.CourseID, assignment.SubmissionID, assignment.Publish, assignment.Deadline, assignment.Reviewers, id)
	} else {
		_, err = tx.Exec(query, assignment.Name, assignment.Description, assignment.CourseID, nil, assignment.Publish, assignment.Deadline, assignment.Reviewers, id)
	}
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	query = "UPDATE assignments SET review_id=? WHERE id=?"

	tx, err = db.GetDB().Begin()
	if err != nil {
		return err
	}

	if assignment.ReviewID.Valid {
		_, err = tx.Exec(query, assignment.ReviewID, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err = tx.Exec(query, nil, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

// GetAllFromCourse returns all assignments in a course
func (repo *AssignmentRepository) GetAllFromCourse(courseID int) ([]Assignment, error) {
	result := make([]Assignment, 0)
	query := "SELECT id, name, description, created, publish, deadline, " +
		"course_id, submission_id, review_id FROM assignments WHERE course_id=?"

	rows, err := db.GetDB().Query(query, courseID)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := Assignment{}

		err := rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.Created,
			&temp.Publish, &temp.Deadline, &temp.CourseID, &temp.SubmissionID, &temp.ReviewID)
		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}

	return result, err
}

// GetAnswersFromUser retrieves all answers and their type from an user on a specific assignment
func (repo *AssignmentRepository) GetAnswersFromUser(assignmentID, userID int) ([]Answer, error) {
	result := make([]Answer, 0)
	query := "SELECT type, answer FROM user_submissions WHERE user_id=? AND assignment_id=?"
	rows, err := db.GetDB().Query(query, userID, assignmentID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var temp Answer

		err := rows.Scan(&temp.Type, &temp.Value)
		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}

	return result, err
}

// HasUserSubmitted checks if a user have submitted a submission to an assignment
func (repo *AssignmentRepository) HasUserSubmitted(assignmentID, userID int) (bool, error) {
	query := "SELECT COUNT(id) FROM user_submissions WHERE user_id=? AND assignment_id=?"
	rows, err := db.GetDB().Query(query, userID, assignmentID)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var temp int

		err := rows.Scan(&temp)
		if err != nil {
			return false, err
		}

		if temp == 0 {
			return false, err
		}
	}

	return true, err
}

// HasReview checks if the assignments has a review form
func (repo *AssignmentRepository) HasReview(id int) (bool, error) {
	var result sql.NullInt64

	query := "SELECT review_id FROM assignments WHERE id=?"
	rows, err := db.GetDB().Query(query, id)
	if err != nil {
		return result.Valid, err
	}

	for rows.Next() {
		err := rows.Scan(&result)
		return result.Valid, err
	}

	return result.Valid, err
}

// HasAutoValidation checks if the assignment has auto validation
func (repo *AssignmentRepository) HasAutoValidation(id int) (bool, error) {
	var result sql.NullInt64

	query := "SELECT validation_id FROM assignments WHERE id=?"
	rows, err := db.GetDB().Query(query, id)
	if err != nil {
		return result.Valid, err
	}

	for rows.Next() {
		err := rows.Scan(&result)
		return result.Valid, err
	}

	return result.Valid, err
}
