package model

import "time"

// ForgottenPass struct for keeping the data for forgotten password
type ForgottenPass struct {
	ID        int       `jsoh:"id"`
	Hash      string    `json:"hash"`
	UserID    int       `json:"userid"`
	TimeStamp time.Time `json:"timestamp"`
}
