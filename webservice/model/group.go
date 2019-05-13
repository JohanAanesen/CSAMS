package model

// Group struct
type Group struct {
	ID           int    `json:"id"`
	AssignmentID int    `json:"assignment_id"`
	Name         string `json:"name"`
	Users        []User `json:"users"`
	Creator      int    `json:"creator"`
}
