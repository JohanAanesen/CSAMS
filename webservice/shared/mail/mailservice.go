package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// SendMail sends mail to the user with mailservice
func (mail Mail) SendMail(email string) error {

	// Make a temporary struct for posting to mailservice
	jsonData := struct {
		Authentication string `json:"authentication"`
		Email          string `json:"email"`
		Link           string `json:"link"`
	}{
		Authentication: os.Getenv("MAIL_AUTH"),
		Email:          email,
		Link:           "https://www.google.no/",
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
