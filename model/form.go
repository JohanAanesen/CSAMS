package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"time"
)

type Form struct {
	ID          int       `json:"id" db:"id"`
	Prefix      string    `json:"prefix" db:"prefix"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`

	Fields      []struct {
		ID          int      `json:"id" db:"id"`
		Type        string   `json:"type" db:"type"`
		Name        string   `json:"name" db:"name"`
		Label       string   `json:"label" db:"label"`
		Description string   `json:"description" db:"description"`
		Order       int      `json:"order" db:"order"`
		Weight      int      `json:"weight" db:"weight"`
		Choices     []string `json:"choices" db:"choices"`
	} `json:"fields"`
}

type FormRepository struct {

}

func (repo *FormRepository) Insert(form Form) (int64, error) {
	// Insertions Query
	query := "INSERT INTO forms (prefix, name, description) VALUES (?, ?, ?);"
	// Execute query with parameters
	rows, err := db.GetDB().Exec(query, form.Prefix, form.Name, form.Description)
	// Check for error
	if err != nil {
		return -1, err
	}

	// Get last inserted id from table
	id, err := rows.LastInsertId()
	// Check for error
	if err != nil {
		return -1, err
	}

	return id, nil
}