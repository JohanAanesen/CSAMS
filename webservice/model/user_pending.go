package model

import "database/sql"

// UserRegistrationPending struct for keeping data for the table user_pending
type UserRegistrationPending struct {
	ID           int            `json:"id"`
	Name         sql.NullString `json:"name"`
	Email        string         `json:"email"`
	Password     sql.NullString `json:"password"`
	ValidationID int            `json:"validation_id"`
}
