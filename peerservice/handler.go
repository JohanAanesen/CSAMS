package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct{
	Submissions Submissions
	AfterShuffle Submissions
}

// HandlerGET
func HandlerGET(w http.ResponseWriter, r *http.Request) {

	payload := Payload{
		Submissions: GetSubmissions(1),
		AfterShuffle: GetSubmissions(1).shuffle(),
	}


	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}

}

// HandlerPOST
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Test :)")

}