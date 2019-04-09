package model

// UserRegistrationPending struct for keeping data for the table user_pending
type UserRegistrationPending struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	EmailStudent string `json:"email_student"`
	Password     string `json:"password"`
	ValidationID int    `json:"validation_id"`
}
