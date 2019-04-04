package model

// UserPending struct for keeping user_pending
type UserPending struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	EmailStudent string `json:"email_student"`
	Password     string `json:"password"`
	ValidationID int    `json:"validation_id"`
}
