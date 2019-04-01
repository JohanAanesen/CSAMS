package main

import (
	"log"
	"net/http"
	"net/smtp"
)

// HandlerGET handles GET requests
func HandlerGET(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
	// SendMail([]string{"RECEIVERMAIL"}, "MESSAGE", "SUBJECT")
}

// HandlerPOST handles POST requests
func HandlerPOST(w http.ResponseWriter, r *http.Request) {
	//SendEmail(w, r)
}

// SendMail sends mail to people
func SendMail(to []string, message string, subject string) {

	auth := smtp.PlainAuth("", "SENDEREMAIL", "PASSWORD", "smtp.gmail.com")

	msg := []byte("To: " + to[0] + "\nSubject: " + subject + "\n" + message)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "SENDEREMAIL", to, msg)
	if err != nil {
		log.Println(err.Error())
	}
}
