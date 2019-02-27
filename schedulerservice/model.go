package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//Payload struct
//https://stackoverflow.com/questions/28254102/how-to-unmarshal-json-into-interface-in-go
type Payload struct {
	ScheduledTime time.Time       `json:"scheduled_time"`
	Task          string          `json:"task"`
	Data          json.RawMessage
}

type Task interface {
	triggerTask()
}

//PeerTask struct
type PeerTask struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
	Reviewers      int    `json:"reviewers"`
}

func (peer PeerTask) triggerTask() {

}

func schedulePeerTask(payload Payload)bool{
	var peerTask PeerTask
	if err := json.Unmarshal([]byte(payload.Data), &peerTask); err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(peerTask.Authentication, peerTask.Reviewers, peerTask.SubmissionID)
	return true
}