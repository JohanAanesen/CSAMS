package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"log"
	"time"
)

// AssignmentRepository struct
type AssignmentRepository struct {
	db *sql.DB
}

// NewAssignmentRepository func
func NewAssignmentRepository(db *sql.DB) *AssignmentRepository {
	return &AssignmentRepository{
		db: db,
	}
}

// Fetch func
func (repo *AssignmentRepository) Fetch(id int) (*model.Assignment, error) {
	/*
		loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
		if err != nil {
			log.Println(err.Error())
		}
	*/
	result := model.Assignment{}

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, review_deadline, reviewers, validation_id FROM assignments WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		var reviewDeadline sql.NullString

		err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.Created,
			&result.Publish, &result.Deadline, &result.CourseID, &result.SubmissionID,
			&result.ReviewID, &reviewDeadline, &result.Reviewers, &result.ValidationID)

		if err != nil {
			return &result, err
		}

		if reviewDeadline.Valid {
			result.ReviewDeadline, err = time.Parse("2006-01-02T15:04:05Z", reviewDeadline.String)
			if err != nil {
				log.Println("review deadline time.Parse error:", err)
			}
		}
	}

	return &result, err
}

// FetchAll func
func (repo *AssignmentRepository) FetchAll() ([]*model.Assignment, error) {
	result := make([]*model.Assignment, 0)

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, review_deadline, reviewers, validation_id FROM assignments"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Assignment{}
		var reviewDeadline sql.NullString

		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.Created,
			&temp.Publish, &temp.Deadline, &temp.CourseID, &temp.SubmissionID,
			&temp.ReviewID, &reviewDeadline, &temp.Reviewers, &temp.ValidationID)
		if err != nil {
			return result, err
		}

		if reviewDeadline.Valid {
			temp.ReviewDeadline, err = time.Parse("2006-01-02T15:04:05Z", reviewDeadline.String)
			if err != nil {
				log.Println("review deadline time.Parse error:", err)
			}
		}

		result = append(result, &temp)
	}

	return result, err
}

// Insert func
func (repo *AssignmentRepository) Insert(assignment model.Assignment) (int, error) {
	var id int64

	query := "INSERT INTO assignments (name, description, created, publish, deadline, course_id) VALUES (?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	created := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	rows, err := tx.Exec(query, assignment.Name, assignment.Description, created,
		assignment.Publish, assignment.Deadline, assignment.CourseID)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	if assignment.SubmissionID.Valid {
		query := "UPDATE assignments SET submission_id = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.SubmissionID, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}

	if assignment.ReviewID.Valid {
		query := "UPDATE assignments SET review_id = ?, review_deadline = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.ReviewID, assignment.ReviewDeadline, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}

	if assignment.Reviewers.Valid {
		query := "UPDATE assignments SET reviewers = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.Reviewers, id)
		if err != nil {
			tx.Rollback()
			return int(id), err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}

// Update func
func (repo *AssignmentRepository) Update(assignment model.Assignment) error {
	query := "UPDATE assignments SET name = ?, description = ?, publish = ?, deadline = ?, course_id = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, assignment.Name, assignment.Description,
		assignment.Publish, assignment.Deadline, assignment.CourseID, assignment.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if assignment.SubmissionID.Valid {
		query := "UPDATE assignments SET submission_id = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.SubmissionID, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if assignment.ReviewID.Valid {
		query := "UPDATE assignments SET review_id = ?, review_deadline = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.ReviewID, assignment.ReviewDeadline, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if assignment.Reviewers.Valid {
		query := "UPDATE assignments SET reviewers = ? WHERE id = ?"
		_, err := tx.Exec(query, assignment.Reviewers, assignment.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}
