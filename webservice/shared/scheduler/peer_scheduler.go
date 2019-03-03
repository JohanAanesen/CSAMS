package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//PeerTask struct
type PeerTask struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
	Reviewers      int    `json:"reviewers"`
}



var dummyUpdate = struct {
	Authentication string    `json:"authentication"`
	SubmissionID   int       `json:"submission_id"`
	ScheduledTime  time.Time `json:"scheduled_time"`
}{
	Authentication: os.Getenv("PEER_AUTH"),
	SubmissionID:   1,
	ScheduledTime:  time.Now().Add(time.Hour * 2351467),
}

var dummyDelete = struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
}{
	Authentication: os.Getenv("PEER_AUTH"),
	SubmissionID:   1,
}

func (scheduler Scheduler)SchedulePeerReview(subID int, reviewers int, scheduledTime time.Time)error{
	// PeerTask, this is what is being sent to the peerservice

	var peerTask = PeerTask{
		Authentication: os.Getenv("PEER_AUTH"),
		SubmissionID:subID,
		Reviewers:reviewers,
	}

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"scheduled_time": scheduledTime,
		"task": "peer",
		"submission_id": subID,
		"data": peerTask,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil{
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

func (scheduler Scheduler)UpdateSchedule(subID int, reviewers int, scheduledTime time.Time)error{
	// PeerTask, this is what is being sent to the peerservice

	var peerTask = PeerTask{
		Authentication: os.Getenv("PEER_AUTH"),
		SubmissionID:subID,
		Reviewers:reviewers,
	}

	//this is what is being sent to the scheduler service
	jsonData := map[string]interface{}{
		"scheduled_time": scheduledTime,
		"task": "peer",
		"submission_id": subID,
		"data": peerTask,
	}

	//this is just sending the request
	jsonValue, err := json.Marshal(jsonData)
	if err != nil{
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