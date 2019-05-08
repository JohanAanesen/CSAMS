package model

import (
	"database/sql"
	"time"
)

// ValidationEmail struct for keeping the data for the validation table for confirming email address, forgotten password and adding secondary email
type ValidationEmail struct {
	ID        int           `json:"id"`
	Hash      string        `json:"hash"`
	UserID    sql.NullInt64 `json:"userid"`
	Valid     bool          `json:"valid"`
	TimeStamp time.Time     `json:"timestamp"`
}
