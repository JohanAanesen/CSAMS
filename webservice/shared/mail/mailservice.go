package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// MailForgottenPassword sends mail to the user with mailservice
func (mail Mail) MailForgottenPassword(email string, link string) error {

	// Make a temporary struct for posting to mailservice
	jsonData := struct {
		Authentication string `json:"authentication"`
		Email          string `json:"email"`
		Link           string `json:"link"`
	}{
		Authentication: os.Getenv("MAIL_AUTH"),
		Email:          email,
		Link:           link,
	}

	// This is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8085" // mailservice

	if os.Getenv("MAIL_SERVICE") != "" {
		url = "http://" + os.Getenv("MAIL_SERVICE") // mailservice address changed in env var
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}

	return nil
}

func (mail Mail) SendSingleRecipient(email string, subject string, message string) error {

	// Make a temporary struct for posting to mailservice
	jsonData := struct {
		Authentication string `json:"authentication"`
		Email          string `json:"email"`
		Subject        string `json:"subject"`
		Message        string `json:"message"`
	}{
		Authentication: os.Getenv("MAIL_AUTH"),
		Email:          email,
		Subject:        subject,
		Message:        message,
	}

	// This is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8085/single" // mailservice

	if os.Getenv("MAIL_SERVICE") != "" {
		url = "http://" + os.Getenv("MAIL_SERVICE") + "/single" // mailservice address changed in env var
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}

	return nil
}

func (mail Mail) SendMultipleRecipient(emails []string, subject string, message string) error {

	// Make a temporary struct for posting to mailservice
	jsonData := struct {
		Authentication string   `json:"authentication"`
		Emails         []string `json:"emails"`
		Subject        string   `json:"subject"`
		Message        string   `json:"message"`
	}{
		Authentication: os.Getenv("MAIL_AUTH"),
		Emails:         emails,
		Subject:        subject,
		Message:        message,
	}

	// This is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8085/multiple" // mailservice

	if os.Getenv("MAIL_SERVICE") != "" {
		url = "http://" + os.Getenv("MAIL_SERVICE") + "/multiple" // mailservice address changed in env var
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}

	return nil
}
