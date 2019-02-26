package model

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	_ "github.com/go-sql-driver/mysql" //database driver
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
}

// AssignmentRepository holds all assignments, and DB-functions
type AssignmentRepository struct{}

// GetSingle retrieves a single Assignment based on Primary Key (id)
func (repo *AssignmentRepository) GetSingle(id int) (Assignment, error) {
	// Declare empty struct
	var result Assignment

	// Create query string
	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id FROM assignments WHERE id = ?"
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
			&result.SubmissionID, &result.ReviewID)
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

	// Create query string
	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id FROM assignments;"
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

	// Create query string
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
func (repo *AssignmentRepository) Insert(assignment Assignment) error {
	// Create query string
	query := "INSERT INTO assignments (name, description, publish, deadline, course_id) VALUES (?, ?, ?, ?, ?);"
	// Prepare and execute query
	rows, err := db.GetDB().Exec(query, assignment.Name, assignment.Description, assignment.Publish, assignment.Deadline, assignment.CourseID)
	// Check for error
	if err != nil {
		return err
	}

	// Get last inserted ID
	id, err := rows.LastInsertId()
	// Check for error
	if err != nil {
		return err
	}

	// Check if we have set a submission_id
	if assignment.SubmissionID.Valid {
		// Create query string
		query := "UPDATE assignments SET submission_id = ? WHERE id = ?;"
		// Prepare and execute query
		rows, err := db.GetDB().Query(query, assignment.SubmissionID, id)
		// Check for error
		if err != nil {
			return err
		}
		// Close connection
		defer rows.Close()
	}

	// Check if we have set a review_id
	if assignment.ReviewID.Valid {
		// Create query string
		query := "UPDATE assignments SET review_id = ? WHERE id = ?;"
		// Prepare and execute query
		rows, err := db.GetDB().Query(query, assignment.ReviewID, id)
		// Check for error
		if err != nil {
			return err
		}
		// Close connection
		defer rows.Close()
	}

	return nil
}

// Update an assignment based on the ID and the data inside an Assignment-object
func (repo *AssignmentRepository) Update(id int, assignment Assignment) error {
	query := "UPDATE assignments SET name=?, description=?, course_id=?, publish=?, deadline=? WHERE id=?"

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, assignment.Name, assignment.Description, assignment.CourseID, assignment.Publish, assignment.Deadline, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if assignment.SubmissionID.Valid {
		query = "UPDATE assignments SET submission_id=? WHERE id=?"
		tx, err = db.GetDB().Begin()
		if err != nil {
			return err
		}

		_, err = tx.Exec(query, assignment.SubmissionID, id)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	if assignment.ReviewID.Valid {
		query = "UPDATE assignments SET review_id=? WHERE id=?"
		tx, err = db.GetDB().Begin()
		if err != nil {
			return err
		}

		_, err = tx.Exec(query, assignment.ReviewID, id)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	return err
}

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