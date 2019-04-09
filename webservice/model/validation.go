package model

import "time"

// ValidationEmail struct for keeping the data for the validation table for confirming email address and forgotten password
type ValidationEmail struct {
	ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	UserID    int       `json:"userid"`
	Valid     bool      `json:"valid"`
	TimeStamp time.Time `json:"timestamp"`
}
