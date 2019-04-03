package model

import "time"

// Validation struct for keeping the data for forgotten password
type Validation struct {
	ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	UserID    int       `json:"userid"`
	Valid     bool      `json:"valid"`
	TimeStamp time.Time `json:"timestamp"`
}
