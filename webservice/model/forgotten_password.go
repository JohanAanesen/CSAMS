package model

// ForgottenPass struct for keeping the data for forgotten password
type ForgottenPass struct {
	Hash      string `json:"hash"`
	UserID    int    `json:"userid"`
	TimeStamp string `json:"timestamp"`
}
