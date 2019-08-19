package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// ForgottenPassGET handles GET requests
func ForgottenPassGET(w http.ResponseWriter, r *http.Request) {
	log.Println("This API only accepts POST-requests")
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// ForgottenPassPOST handles POST requests
func ForgottenPassPOST(w http.ResponseWriter, r *http.Request) {

	// Check that request body is not empty
	if r.Body == nil {
		log.Println("No Body in request")
		http.Error(w, "No Body in request", http.StatusBadRequest)
		return
	}

	payload := ForgottenPasswordMail{}

	// Decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error())
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}

	// Close body
	err = r.Body.Close()
	if err != nil {
		log.Println("Something went wrong closing body" + err.Error())
		http.Error(w, "Something went wrong closing body", http.StatusInternalServerError)
		return
	}

	// Authenticate request
	if payload.Authentication != os.Getenv("MAIL_AUTH") { // Don't accept requests from places we don't know
		log.Println("Unauthorized request")
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	// Send forgotten password link
	err = SendMailForgottenPassword(payload.Email, payload.Link)
	if err != nil {
		log.Println("Send Forgotten Password Mail", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SingleMailGET handles GET requests
func SingleMailGET(w http.ResponseWriter, r *http.Request) {
	log.Println("This API only accepts POST-requests")
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// SingleMailPOST handles POST requests for sending an email to an single recipient
func SingleMailPOST(w http.ResponseWriter, r *http.Request) {
	// Check that request body is not empty
	if r.Body == nil {
		log.Println("No Body in request")
		http.Error(w, "No Body in request", http.StatusBadRequest)
		return
	}

	singleRes := SingleReceiver{}

	// Decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&singleRes)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error())
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}

	// Close body
	err = r.Body.Close()
	if err != nil {
		log.Println("Something went wrong closing body" + err.Error())
		http.Error(w, "Something went wrong closing body", http.StatusInternalServerError)
		return
	}

	// Authenticate request
	if singleRes.Authentication != os.Getenv("MAIL_AUTH") { // Don't accept requests from places we don't know
		log.Println("Unauthorized request")
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	// Send forgotten password link
	err = SendMailSingleRecipient(singleRes)
	if err != nil {
		log.Println("Send Single Mail", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// MultipleMailGET handles GET requests
func MultipleMailGET(w http.ResponseWriter, r *http.Request) {
	log.Println("This API only accepts POST-requests")
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// MultipleMailPOST handles POST request for sending an mail to multiple receivers
func MultipleMailPOST(w http.ResponseWriter, r *http.Request) {
	// Check that request body is not empty
	if r.Body == nil {
		log.Println("No Body in request")
		http.Error(w, "No Body in request", http.StatusBadRequest)
		return
	}

	receivers := MultipleReceivers{}

	// Decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&receivers)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error())
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}

	// Close body
	err = r.Body.Close()
	if err != nil {
		log.Println("Something went wrong closing body" + err.Error())
		http.Error(w, "Something went wrong closing body", http.StatusInternalServerError)
		return
	}

	// Authenticate request
	if receivers.Authentication != os.Getenv("MAIL_AUTH") { // Don't accept requests from places we don't know
		log.Println("Unauthorized request")
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	// Send forgotten password link
	err = SendMailMultipleRecipients(receivers)
	if err != nil {
		log.Println("Send Multiple Mails", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// SendMailForgottenPassword sends a link to the user to change password
func SendMailForgottenPassword(recipient string, link string) error {

	// Convert to string array
	userEmail := strings.Fields(recipient)
	subject := "Forgotten Password | CSAMS"
	message := "Hi\n" +
		"This email is sent by the CS Assignment Submission System.\n" +
		"There has been requested a password recovery for this account on CSAMS\n" +
		"If you have not requested this and suspect a hacking attempt, please contact your lecturer.\n\n" +
		"Click this link to reset yout password:\n" +
		link

	err := sendMail("To", userEmail, subject, message)
	if err != nil {
		return err
	}

	return nil
}

// SendMailSingleRecipient sends a message to a single recipient
func SendMailSingleRecipient(receiver SingleReceiver) error {
	// Convert to string array
	userEmail := strings.Fields(receiver.Email)

	err := sendMail("To", userEmail, receiver.Subject, receiver.Message)
	if err != nil {
		return err
	}

	return nil
}

// SendMailMultipleRecipients sends a mail to multiple recipients
func SendMailMultipleRecipients(receivers MultipleReceivers) error {

	err := sendMail("Bbc", receivers.Emails, receivers.Subject, receivers.Message)
	if err != nil {
		return err
	}

	return nil
}

// sendMail sends the mail to recipient(s)
func sendMail(toType string, recipients []string, subject string, message string) error {

	users := strings.Join(recipients, ",")

	// Get authentication
	auth := smtp.PlainAuth("", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("MAILPROVIDER"))

	// Write message
	msg := []byte(toType + ": " + users + "\nSubject:" + subject + "\n" + message)

	// Send mail and check for errors
	err := smtp.SendMail(os.Getenv("MAILPROVIDER")+":587", auth, os.Getenv("USERNAME"), recipients, msg)
	if err != nil {
		return err
	}

	return nil
}
