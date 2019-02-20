package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerGET
func HandlerGET(w http.ResponseWriter, r *http.Request) {

	submissions := GetSubmissions(1)

	if err := json.NewEncoder(w).Encode(submissions); err != nil {
		panic(err)
	}

}

// HandlerPOST
func HandlerPOST(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Test :)")

}