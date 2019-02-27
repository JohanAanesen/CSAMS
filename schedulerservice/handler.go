package main

import (
	"net/http"
)

//Payload struct
type Payload struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submissionid"`
	Reviewers      int    `json:"reviewers"`
}

// PeerGET handles GET requests
func PeerGET(w http.ResponseWriter, r *http.Request) {

}

// PeerPOST handles POST requests
func PeerPOST(w http.ResponseWriter, r *http.Request) {

}

// PeerPUT handles PUT requests
func PeerPUT(w http.ResponseWriter, r *http.Request) {

}

// PeerDELETE handles DELETE requests
func PeerDELETE(w http.ResponseWriter, r *http.Request) {

}