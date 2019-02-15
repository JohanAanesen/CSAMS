package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"time"
)

// Assignment hold the data for a single assignment
type Assignment struct {
	ID           int       `json:"id" db:"ID"`
	CourseID     int       `json:"course_id" db:"assignment_course_ID"`
	Title        string    `json:"title" db:"assignment_title"`
	Description  string    `json:"description" db:"assignment_description"`
	Publish      time.Time `json:"publish" db:"assignment_publish"`
	Deadline     time.Time `json:"deadline" db:"assignment_deadline"`
	EnableReview bool      `json:"enable_review" db:"assignment_enable_review"`
	Type         string    `json:"type" db:"assignment_type"`
}

// AssignmentDatabase holds all assignments, and DB-functions
type AssignmentDatabase struct{}

func (adb *AssignmentDatabase) GetAll() []Assignment {
	return nil
}

func (adb *AssignmentDatabase) Insert(a Assignment) (bool, error) {
	rows, err := db.GetDB().Query("INSERT INTO assignments (courseid, assignment_title, assignment_description, assignment_publish, assignment_deadline, assignment_review)"+
		"VALUES (?, ?, ?, ?, ?, ?)", a.CourseID, a.Title, a.Description, a.Publish, a.Deadline, a.EnableReview)
	defer rows.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}
