package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//PeerTask struct
type PeerTask struct {
	Authentication string `json:"authentication"`
	AssignmentID   int    `json:"assignment_id"`
	Reviewers      int    `json:"reviewers"`
}

// SchedulePeerReview schedules a peer review task with scheduler service
func (scheduler Scheduler) SchedulePeerReview(assID int, reviewers int, scheduledTime time.Time) error {
	// PeerTask, this is what is being sent to the peerservice

	var peerTask = PeerTask{
		Authentication: os.Getenv("PEER_AUTH"),
		AssignmentID:   assID,
		Reviewers:      reviewers,
	}

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"scheduled_time": scheduledTime,
		"task":           "peer",
		"assignment_id":  assID,
		"data":           peerTask,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8086" //schedulerservice

	if os.Getenv("SCHEDULE_SERVICE") != "" {
		url = "http://" + os.Getenv("SCHEDULE_SERVICE") //schedulerservice address changed in env var
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}

	return nil
}

// UpdateSchedule updates a schedule task on service
func (scheduler Scheduler) UpdateSchedule(assID int, reviewers int, scheduledTime time.Time) error {

	var peerTask = PeerTask{
		Authentication: os.Getenv("PEER_AUTH"),
		AssignmentID:   assID,
		Reviewers:      reviewers,
	}

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"scheduled_time": scheduledTime,
		"assignment_id":  assID,
		"data":           peerTask,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8086" //schedulerservice

	if os.Getenv("SCHEDULE_SERVICE") != "" {
		url = "http://" + os.Getenv("SCHEDULE_SERVICE") //schedulerservice address changed in env var
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := client.Do(req) //run PUT request
	if err != nil {
		// handle error
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	return nil
}

// DeleteSchedule deletes a planned task from schedule service
func (scheduler Scheduler) DeleteSchedule(assID int) error {

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"assignment_id":  assID,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	url := "http://localhost:8086" //schedulerservice

	if os.Getenv("SCHEDULE_SERVICE") != "" {
		url = "http://" + os.Getenv("SCHEDULE_SERVICE") //schedulerservice address changed in env var
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := client.Do(req) //run PUT request
	if err != nil {
		// handle error
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	return nil
}

//SchedulerExists returns true if a scheduler with subID and assID identical exists
func (scheduler Scheduler) SchedulerExists(assID int) bool {

	parameters := fmt.Sprintf("/%v", assID)

	url := "http://localhost:8086" //schedulerservice

	if os.Getenv("SCHEDULE_SERVICE") != "" {
		url = "http://" + os.Getenv("SCHEDULE_SERVICE") //schedulerservice address changed in env var
	}

	response, err := http.Get(url + parameters)
	if err != nil {
		log.Printf("The HTTP request to schedulerservice failed with error %s\n", err)
		return false
	}

	if response.StatusCode == 404 {
		return false
	}

	return true
}
