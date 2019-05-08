package model

import (
	"database/sql"
)

//User struct to hold session data
type User struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	EmailStudent  string         `json:"emailstudent"`
	EmailPrivate  sql.NullString `json:"emailprivate,omitempty"`
	Teacher       bool           `json:"teacher"`
	Authenticated bool           `json:"authenticated"`
}
