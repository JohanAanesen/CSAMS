package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"log"
	"time"
)

// AssignmentRepository represents the layer between the database and the Assignment-model
type AssignmentRepository struct {
	db *sql.DB
}

// NewAssignmentRepository returns a pointer to a new AssignmentRepository
func NewAssignmentRepository(db *sql.DB) *AssignmentRepository {
	return &AssignmentRepository{
		db: db,
	}
}

// Fetch a single assignment from the database
func (repo *AssignmentRepository) Fetch(id int) (*model.Assignment, error) {
	/*
		loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
		if err != nil {
			log.Println(err.Error())
		}
	*/
	// Initialize an empty assignment
	result := model.Assignment{}

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_enabled, review_id, review_deadline, reviewers, validation_id FROM assignments WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	// Close connection last
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Create empty null-string
		var reviewDeadline sql.NullString
		// Scan columns from the row
		err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.Created,
			&result.Publish, &result.Deadline, &result.CourseID, &result.SubmissionID, &result.ReviewEnabled,
			&result.ReviewID, &reviewDeadline, &result.Reviewers, &result.ValidationID)

		if err != nil {
			return nil, err
		}
		// Check if review deadline return is valid
		if reviewDeadline.Valid {
			// Parse string
			result.ReviewDeadline, err = time.Parse("2006-01-02T15:04:05Z", reviewDeadline.String)
			if err != nil {
				log.Println("review deadline time.Parse error:", err)
			}
		}
	}
	// Return result
	return &result, nil
}

// FetchAll all assignment from the database
func (repo *AssignmentRepository) FetchAll() ([]*model.Assignment, error) {
	// Create empty assignment slice
	result := make([]*model.Assignment, 0)

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_enabled, review_id, review_deadline, reviewers, validation_id FROM assignments"

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Temporary assignment object
		temp := model.Assignment{}
		var reviewDeadline sql.NullString
		// Scan all columns from row
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.Created,
			&temp.Publish, &temp.Deadline, &temp.CourseID, &temp.SubmissionID, &temp.ReviewEnabled,
			&temp.ReviewID, &reviewDeadline, &temp.Reviewers, &temp.ValidationID)
		if err != nil {
			return nil, err
		}
		// Check if review deadline string is valid
		if reviewDeadline.Valid {
			// Parse time
			temp.ReviewDeadline, err = time.Parse("2006-01-02T15:04:05Z", reviewDeadline.String)
			if err != nil {
				log.Println("review deadline time.Parse error:", err)
			}
		}
		// Append temporary assignment to result
		result = append(result, &temp)
	}
	// Return result
	return result, nil
}

// Insert an assignment into the database
func (repo *AssignmentRepository) Insert(assignment model.Assignment) (int, error) {
	// Integer to hold the id of last inserted row
	var id int64
	query := "INSERT INTO assignments (name, description, created, publish, deadline, course_id, review_enabled) VALUES (?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}
	// Create timestamp-string from current time
	created := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())
	// Execute transaction
	rows, err := tx.Exec(query, assignment.Name, assignment.Description, created,
		assignment.Publish, assignment.Deadline, assignment.CourseID, assignment.ReviewEnabled)
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
	// Check if submission ID is valid
	if assignment.SubmissionID.Valid {
		// Query string
		query := "UPDATE assignments SET submission_id = ?, group_delivery = ? WHERE id = ?"
		// Set group delivery to zero, and to 1 if is true
		groupDelivery := 0
		if assignment.GroupDelivery {
			groupDelivery = 1
		}
		// Execute query
		_, err := tx.Exec(query, assignment.SubmissionID, groupDelivery, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}
	// Check if review ID is valid
	if assignment.ReviewID.Valid {
		// Query string
		query := "UPDATE assignments SET review_id = ?, review_deadline = ? WHERE id = ?"
		// Execute query
		_, err := tx.Exec(query, assignment.ReviewID, assignment.ReviewDeadline, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}
	// Check if reviewers is valid
	if assignment.Reviewers.Valid {
		// Query string
		query := "UPDATE assignments SET reviewers = ? WHERE id = ?"
		// Execute query
		_, err := tx.Exec(query, assignment.Reviewers, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}
	// Return last inserted id
	return int(id), nil
}

// Update assignment in the database
func (repo *AssignmentRepository) Update(assignment model.Assignment) error {
	query := "UPDATE assignments SET name = ?, description = ?, publish = ?, deadline = ?, course_id = ?, review_enabled = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	// Execute query
	_, err = tx.Exec(query, assignment.Name, assignment.Description,
		assignment.Publish, assignment.Deadline, assignment.CourseID, assignment.ReviewEnabled, assignment.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Check if submission ID is valid
	if assignment.SubmissionID.Valid {
		// Query string
		query := "UPDATE assignments SET submission_id = ?, group_delivery = ? WHERE id = ?"
		// Set group delivery to zero, and to 1 if is true
		groupDelivery := 0
		if assignment.GroupDelivery {
			groupDelivery = 1
		}
		// Execute query
		_, err := tx.Exec(query, assignment.SubmissionID, groupDelivery, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// Check if review ID is valid
	if assignment.ReviewID.Valid {
		// Query string
		query := "UPDATE assignments SET review_id = ?, review_deadline = ? WHERE id = ?"
		// Execute query
		_, err := tx.Exec(query, assignment.ReviewID, assignment.ReviewDeadline, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		query := "UPDATE assignments SET review_id = ? WHERE id = ?"
		_, err := tx.Exec(query, sql.NullInt64{Valid: false}, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// Check if reviewers is valid
	if assignment.Reviewers.Valid {
		// Query string
		query := "UPDATE assignments SET reviewers = ? WHERE id = ?"
		// Execute query
		_, err := tx.Exec(query, assignment.Reviewers, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	// Return nil, all good
	return nil
}
