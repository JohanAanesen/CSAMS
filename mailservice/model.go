package main

type Payload struct {
	Authentication string `json:"authentication"`
	Email          string `json:"email"`
	Link           string `json:"link"`
}
