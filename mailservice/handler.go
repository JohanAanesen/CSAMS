package main

import (
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// HandlerGET handles GET requests
func HandlerGET(w http.ResponseWriter, r *http.Request) {
	log.Println("This API only accepts POST-requests")
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// HandlerPOST handles POST requests
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	// Check that request body is not empty
	if r.Body == nil {
		log.Println("No Body in request")
		http.Error(w, "No Body in request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	link := r.FormValue("link")

	if email == "" {
		log.Println("email value can not be empty!")
		http.Error(w, "email value can not be empty!", http.StatusBadRequest)
		return
	}

	if link == "" {
		log.Println("link value can not be empty!")
		http.Error(w, "link value can not be empty!", http.StatusBadRequest)
		return
	}

	// Send forgotten password link
	err := SendMailForgottenPassword(email, link)
	if err != nil {
		log.Println("Send Mail", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SendMailForgottenPassword sends a link to the user to change password
func SendMailForgottenPassword(recipient string, link string) error {

	// Convert to string array
	userEmail := strings.Fields(recipient)

	// Get authentication
	auth := smtp.PlainAuth("", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "smtp.gmail.com")

	// Write message
	msg := []byte("To: " + recipient + "\nSubject: Forgotten Password | CS Assignments\n" +
		"Change password\n" +
		"Looks like you have forgotten your password.\n" +
		"If this was not you, please disregard this email.\n\n" +
		"Click this link to reset your password:\n" +
		link +
		"\nGood luck :)")

	// Send mail and check for errors
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("USERNAME"), userEmail, msg)
	if err != nil {
		return err
	}

	return nil
}
