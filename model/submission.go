package model

type Submission struct {
	ID     int   `json:"id" db:"id"`
	FormID int   `json:"-" db:"form_id"`
	Form   *Form `json:"form"`
}
