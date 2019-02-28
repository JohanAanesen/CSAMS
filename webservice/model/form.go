package model

import (
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"time"
)

// Form TODO (Svein): comment
type Form struct {
	ID          int       `json:"id" db:"id"`
	Prefix      string    `json:"prefix" db:"prefix"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`
	Fields      []Field   `json:"fields"`
}

// Field TODO (Svein): comment
type Field struct {
	ID          int      `json:"id" db:"id"`
	Type        string   `json:"type" db:"type"`
	Name        string   `json:"name" db:"name"`
	Label       string   `json:"label" db:"label"`
	Description string   `json:"description" db:"description"`
	Order       int      `json:"order" db:"priority"`
	Weight      int      `json:"weight" db:"weight"`
	Choices     []string `json:"choices" db:"choices"`
}

// Answer2 struct used for storing answers from users in forms
type Answer struct {
	Type  string
	Value string
}

// FormRepository ... TODO (Svein): comment
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

// Get a single form based on the Primary Key, 'id'
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
	}

	return Form{}, errors.New("form: Could not do rows.Next()")
}

// GetFromAssignmentID get form from the assignment id key
func (repo *FormRepository) GetFromAssignmentID(assignmentID int) (Form, error) {

	// Create query-string
	query := "SELECT f.form_id, f.id, f.type, f.name, f.label, f.description, f.priority, f.weight from fields AS f WHERE f.form_id IN (SELECT s.form_id FROM submissions AS s WHERE id IN (SELECT a.submission_id FROM assignments AS a WHERE id=?)) ORDER BY f.priority"

	// Perform query
	rows, err := db.GetDB().Query(query, assignmentID)

	// Declare an empty Form
	form := Form{}

	// Check for error
	if err != nil {
		return form, err
	}

	// NOT 'if rows.Next()'!! THAT IS NOT THE SAME FUCK!
	for rows.Next() {
		var formID int
		var fieldID int
		var fieldType string
		var name string
		var label string
		var desc string
		var priority int
		var weight int

		// Scan
		err = rows.Scan(&formID, &fieldID, &fieldType, &name, &label, &desc, &priority, &weight) //, &choices)
		// Check for error
		if err != nil {
			return form, err
		}

		// This only needs to be set one time really :/
		//form.ID = formID

		form.Fields = append(form.Fields, Field{
			ID:          formID,
			Type:        fieldType,
			Name:        name,
			Label:       label,
			Description: desc,
			Order:       priority,
			Weight:      weight,
			//Choices:     choices, // TODO : uncomment this
		})
	}

	// TODO brede use sql.null<type>

	return form, nil
}
