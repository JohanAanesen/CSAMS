package model

import "time"

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

// SELECT * FROM fields WHERE form_id = 1 ORDER BY fields.order
