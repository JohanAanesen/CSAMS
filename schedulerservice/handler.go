package main

import (
	"encoding/json"
	"github.com/JohanAanesen/CSAMS/schedulerservice/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type response struct {
	Success bool `json:"success"`
}

// IndexGET handles GET requests
func IndexGET(w http.ResponseWriter, r *http.Request) {
	//todo
	http.Header.Add(w.Header(), "content-type", "application/json")
	_ = json.NewEncoder(w).Encode(model.GetPayloads())
}

// IndexSingleGET gets a single schedule from service
func IndexSingleGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assID, err := strconv.Atoi(vars["assid"])
	if err != nil {
		log.Printf("id: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	payload := model.GetPayload(assID)

	if payload.ID != 0 {
		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(payload)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	http.Error(w, "bad request", http.StatusNotFound)
}

// IndexPOST handles POST requests
func IndexPOST(w http.ResponseWriter, r *http.Request) {

	var payload model.Payload

	if r.Body == nil {
		log.Println("No Body in request") //todo real logger
		http.Error(w, "No Body in request", http.StatusBadRequest)
		return
	}

	//decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error()) //todo real logger
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//authenticate
	if payload.Authentication != os.Getenv("PEER_AUTH") { //don't accept requests from places we don't know
		log.Println("Unauthorized request") //todo real logger
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	//schedule Task based on type of task
	switch payload.Task {
	case "peer":
		if !model.NewTask(payload) {
			log.Println("Something went wrong decoding request") //todo real logger
			http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
			return
		}
	default:
		//log.Println("Something went wrong decoding request") //todo real logger
		//http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(response{Success: true}); err != nil {
		log.Println("Something went wrong encoding response") //todo real logger
		http.Error(w, "Something went wrong encoding response", http.StatusInternalServerError)
	}
}

// IndexPUT handles PUT requests
func IndexPUT(w http.ResponseWriter, r *http.Request) {
	var update struct {
		Authentication string          `json:"authentication"`
		AssignmentID   int             `json:"assignment_id"`
		ScheduledTime  time.Time       `json:"scheduled_time"`
		Data           json.RawMessage `json:"data"`
	}

	//decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error()) //todo real logger
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//authenticate
	if update.Authentication != os.Getenv("PEER_AUTH") { //don't accept requests from places we don't know
		log.Println("Unauthorized request") //todo real logger
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	model.UpdateTimer(update.Data, update.ScheduledTime, model.GetPayload(update.AssignmentID))

	if err := json.NewEncoder(w).Encode(response{Success: true}); err != nil {
		log.Println("Something went wrong encoding response") //todo real logger
		http.Error(w, "Something went wrong encoding response", http.StatusInternalServerError)
	}
}

// IndexDELETE handles DELETE requests
func IndexDELETE(w http.ResponseWriter, r *http.Request) {
	var delete struct {
		Authentication string `json:"authentication"`
		AssignmentID   int    `json:"assignment_id"`
	}

	//decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&delete)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error()) //todo real logger
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//authenticate
	if delete.Authentication != os.Getenv("PEER_AUTH") { //don't accept requests from places we don't know
		log.Println("Unauthorized request") //todo real logger
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		return
	}

	if !model.DeletePayload(delete.AssignmentID) { //delete the thing
		log.Println("Something wrong deleting timer") //todo real logger
		http.Error(w, "Something wrong deleting timer", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response{Success: true}); err != nil {
		log.Println("Something went wrong encoding response") //todo real logger
		http.Error(w, "Something went wrong encoding response", http.StatusInternalServerError)
	}
}
