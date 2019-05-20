package model

// Notification hold the data for a single notification
type Notification struct {
	ID      int    `json:"id" db:"id"`
	UserID  int    `json:"user_id" db:"user_id"`
	URL     string `json:"url" db:"url"`
	Message string `json:"message" db:"message"`
}
