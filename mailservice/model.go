package main

// ForgottenPasswordMail struct for sending recovery link for forgotten password
type ForgottenPasswordMail struct {
	Authentication string `json:"authentication"`
	Email          string `json:"email"`
	Link           string `json:"link"`
}

// SingleReceiver struct for sending email to one user
type SingleReceiver struct {
	Authentication string `json:"authentication"`
	Email          string `json:"email"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}

// MultipleReceivers struct for sending an email to multiple users
type MultipleReceivers struct {
	Authentication string   `json:"authentication"`
	Emails         []string `json:"emails"`
	Subject        string   `json:"subject"`
	Message        string   `json:"message"`
}
