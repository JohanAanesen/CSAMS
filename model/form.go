package model

import "time"

type Form struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`
	Fields      []struct {
		ID    int    `json:"id" db:"id"`
		Data  string `json:"data" db:"data"`
		Order int    `json:"order" db:"order"`
	} `json:"fields"`
}

// SELECT * FROM fields WHERE form_id = 1 ORDER BY fields.order