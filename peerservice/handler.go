package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
		AfterShuffle: GetSubmissions(1),
	}

	for i := range payload.AfterShuffle.Items {
		j := rand.Intn(i + 1)
		payload.AfterShuffle.Items[i], payload.AfterShuffle.Items[j] = payload.AfterShuffle.Items[j], payload.AfterShuffle.Items[i]
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}

}

// HandlerPOST
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Test :)")

}