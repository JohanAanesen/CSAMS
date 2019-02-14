package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"time"
)

// Assignment hold the data for a single assignment
type Assignment struct {
	ID          int       `json:"id" db:"ID"`
	CourseID    int       `json:"course_id" db:"assignment_course_ID"`
	Title       string    `json:"title" db:"assignment_title"`
	Description string    `json:"description" db:"assignment_description"`
	Publish     time.Time `json:"publish" db:"assignment_publish"`
	Deadline    time.Time `json:"deadline" db:"assignment_deadline"`
}

// AssignmentTable holds all assignments, and DB-functions
type AssignmentTable struct {}

func (table *AssignmentTable) GetAll() []Assignment {
	return nil
}

func (table AssignmentTable) Insert(a Assignment) (bool, error) {
	rows, err := db.GetDB().Query("INSERT INTO assignments (assignment_course_id, assignment_title, assignment_description, assignment_publish, assignment_deadline)" +
		"VALUES (?, ?, ?, ?, ?)")
	defer rows.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}
