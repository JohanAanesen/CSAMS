package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	ShuffledSubmissions   Submissions
	GeneratedPairs Pairs
	Reviewers      int
}

// HandlerGET
func HandlerGET(w http.ResponseWriter, r *http.Request) {

	req := Request{
		SubmissionID: 1,
		Reviewers: 2,
	}

	payload := Payload{
		ShuffledSubmissions: GetSubmissions(1).shuffle(),
		Reviewers: req.Reviewers,
	}

	for i, item := range payload.ShuffledSubmissions.Items{ //iterate all submissions
		for j := 1; j <= req.Reviewers; j++{

			var subpair SubPair

			subpair.UserID = item.UserID
			subpair.SubmissionID = req.SubmissionID

			counter := i+j
			if counter >= len(payload.ShuffledSubmissions.Items){ //if it exceeds array, start from beginning
				counter -= len(payload.ShuffledSubmissions.Items)
			}

			subpair.ReviewID = payload.ShuffledSubmissions.Items[counter].ID

			payload.GeneratedPairs.Items = append(payload.GeneratedPairs.Items, subpair) //save to generated pairs
		}
	}


	savePairs(payload.GeneratedPairs)

/*
	cmd := exec.Command("mysqldump", "-u root --password=root cs53 > backup.sql")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Fprintf(w, "combined out:\n%s\n", string(out))
*/

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}

}

// HandlerPOST
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Test :)")

}
