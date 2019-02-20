package model

import (
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"time"
)

// Form TODO (Svein): comment
type Form struct {
	ID          int       `json:"id" db:"id"`
	Prefix      string    `json:"prefix" db:"prefix"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`

	Fields []struct {
		ID          int      `json:"id" db:"id"`
		Type        string   `json:"type" db:"type"`
		Name        string   `json:"name" db:"name"`
		Label       string   `json:"label" db:"label"`
		Description string   `json:"description" db:"description"`
		Order       int      `json:"order" db:"priority"`
		Weight      int      `json:"weight" db:"weight"`
		Choices     []string `json:"choices" db:"choices"`
	} `json:"fields"`
}

// FormRepository TODO (Svein): comment
type FormRepository struct {
}

// Insert form to database
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

func (repo *FormRepository) Get(id int) (Form, error) {
	// Create query-string
	query := "SELECT id, prefix, name, description, created FROM forms WHERE id = ?"
	// Perform query
	rows, err := db.GetDB().Query(query, id)
	// Check for error
	if err != nil {
		return Form{}, err
	}

	// Check if there is any rows
	if rows.Next() {
		// Declare an empty Form
		var form = Form{}
		// Scan
		err = rows.Scan(&form.ID, &form.Prefix, &form.Name, &form.Description, &form.Created)
		// Check for error
		if err != nil {
			return Form{}, err
		}

		return form, nil
	} else {
		return Form{}, errors.New("form: Could not do rows.Next()")
	}
}
