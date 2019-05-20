package model

type ReviewMessage struct {
	ID           int `json:"id" db:"id"`
	UserReviewer int `json:"user_reviewer" db:"user_reviewer"`
	UserTarget   int `json:"user_target" db:"user_target"`
	ReviewID     int `json:"review_id" db:"review_id"`
	AssignmentID int `json:"assignment_id" db:"assignment_id"`
	Message      int `json:"message" db:"message"`
}
