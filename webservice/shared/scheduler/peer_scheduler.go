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
	SubmissionID   int    `json:"submission_id"`
	AssignmentID   int    `json:"assignment_id"`
	Reviewers      int    `json:"reviewers"`
}

var dummyUpdate = struct {
	Authentication string    `json:"authentication"`
	SubmissionID   int       `json:"submission_id"`
	AssignmentID   int       `json:"assignment_id"`
	ScheduledTime  time.Time `json:"scheduled_time"`
}{
	Authentication: os.Getenv("PEER_AUTH"),
	SubmissionID:   1,
	AssignmentID:   1,
	ScheduledTime:  time.Now().Add(time.Hour * 2351467),
}

var dummyDelete = struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
	AssignmentID   int    `json:"assignment_id"`
}{
	Authentication: os.Getenv("PEER_AUTH"),
	SubmissionID:   1,
	AssignmentID:   1,
}

func (scheduler Scheduler) SchedulePeerReview(subID int, assID int, reviewers int, scheduledTime time.Time) error {
	// PeerTask, this is what is being sent to the peerservice

	var peerTask = PeerTask{
		Authentication: os.Getenv("PEER_AUTH"),
		SubmissionID:   subID,
		Reviewers:      reviewers,
	}

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"scheduled_time": scheduledTime,
		"task":           "peer",
		"submission_id":  subID,
		"assignment_id":  assID,
		"data":           peerTask,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	response, err := http.Post("http://"+os.Getenv("SCHEDULE_SERVICE"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	return nil
}

func (scheduler Scheduler) UpdateSchedule(subID int, assID int, scheduledTime time.Time) error {

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"scheduled_time": scheduledTime,
		"submission_id":  subID,
		"assignment_id":  assID,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://"+os.Getenv("SCHEDULE_SERVICE"), bytes.NewBuffer(jsonValue))
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

func (scheduler Scheduler) DeleteSchedule(subID int, assID int) error {

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"authentication": os.Getenv("PEER_AUTH"),
		"submission_id":  subID,
		"assignment_id":  assID,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, "http://"+os.Getenv("SCHEDULE_SERVICE"), bytes.NewBuffer(jsonValue))
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

func (scheduler Scheduler) SchedulerExists(subID int, assID int) (bool, error) {

	parameters := fmt.Sprintf("?subid=%v&assid=%v", subID, assID)

	response, err := http.Get(os.Getenv("PEER_SERVICE")+parameters)
	if err != nil {
		fmt.Printf("The HTTP request to schedulerservice failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}




	return true, nil
}
