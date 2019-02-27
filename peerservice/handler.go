package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

//Payload struct
type Payload struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submissionid"`
	Reviewers      int    `json:"reviewers"`
	//todo Add date here and then register to schedule service, or directly to schedule service
}

// HandlerGET handles GET requests
func HandlerGET(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This API only accepts POST-requests", http.StatusBadRequest)
	return
}

// HandlerPOST handles POST requests
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	var ShuffledSubmissions Submissions //slice with submissions
	var GeneratedReviewers Pairs        //slice with generated peer reviewers
	var payload Payload                 //payload struct

	//check that request body is not empty
	if r.Body == nil {
		log.Println("Please send a request body")
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	//decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Something went wrong decoding request")
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//authenticate
	if payload.Authentication != os.Getenv("PEER_AUTH") { //don't accept requests from places we don't know
		log.Println("Unauthorized request")
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	//get submissions from database
	ShuffledSubmissions = getSubmissions(payload.SubmissionID).shuffle() //shuffle part important!

	//make sure the nr of reviewers is greater than the number of submissions
	if payload.Reviewers >= len(ShuffledSubmissions) {
		log.Println("Submissions is less than the number of submissions everyone should review.")
		http.Error(w, "Submissions is less than the number of submissions everyone should review.", http.StatusBadRequest)
		return
	}

	//generate peer reviewers from submissions
	for i, item := range ShuffledSubmissions { //iterate all submissions
		for j := 1; j <= payload.Reviewers; j++ {

			var subpair SubPair

			subpair.UserID = item.UserID
			subpair.SubmissionID = payload.SubmissionID

			counter := i + j
			if counter >= len(ShuffledSubmissions) { //if it exceeds array, start from beginning
				counter -= len(ShuffledSubmissions)
			}

			subpair.ReviewID = ShuffledSubmissions[counter].ID

			GeneratedReviewers = append(GeneratedReviewers, subpair) //save to generated pairs
		}
	}

	//store peer reviewers in database
	if !savePairs(GeneratedReviewers) {
		log.Println("Something went wrong storing peer_reviews to database. Try again.")
		http.Error(w, "Something went wrong storing peer_reviews to database. Try again.", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}
}
