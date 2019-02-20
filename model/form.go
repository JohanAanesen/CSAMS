package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"strings"
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

// Insert form to database
func (repo *FormRepository) Insert(form Form) error {
	// Insertions Query
	query := "INSERT INTO forms (prefix, name, description) VALUES (?, ?, ?);"
	// Execute query with parameters
	rows, err := db.GetDB().Exec(query, form.Prefix, form.Name, form.Description)
	// Check for error
	if err != nil {
		return err
	}

	// Get last inserted id from table
	formId, err := rows.LastInsertId()
	// Check for error
	if err != nil {
		return err
	}

	// Loop trough fields in the forms
	for _, field := range form.Fields {
		// Join the array to a single string for 'choices'
		choices := strings.Join(field.Choices, ",")
		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		// Execute the query
		rows, err := db.GetDB().Query(query, int(formId), field.Type, field.Name, field.Label, field.Description, field.Order, field.Weight, choices)
		// Check for error
		if err != nil {
			return err
		}
		// Close the connection
		rows.Close()
	}

	// Return no error
	return nil
}
