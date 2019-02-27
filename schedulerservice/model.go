package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//Timers slice hold all current timers
var Timers  = make(map[int]*time.Timer)

//Payload struct
//https://stackoverflow.com/questions/28254102/how-to-unmarshal-json-into-interface-in-go
type Payload struct {
	ScheduledTime time.Time       `json:"scheduled_time"`
	Task          string          `json:"task"`
	Data          json.RawMessage
}

//Task interface
type Task interface {
	triggerTask()
}

//PeerTask struct
type PeerTask struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
	Reviewers      int    `json:"reviewers"`
}

func init(){

}

func (peer PeerTask) triggerTask() {
	//todo send request to peerservice
}

func schedulePeerTask(payload Payload)bool{
	var peerTask PeerTask
	if err := json.Unmarshal([]byte(payload.Data), &peerTask); err != nil { //read the unread json into peerTask
		log.Fatal(err)
		return false
	}
	fmt.Println(peerTask) //todo remove this

	timeNow := time.Now() //time now
	Duration := payload.ScheduledTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	Timers[peerTask.SubmissionID] = time.NewTimer(Duration) //register timer
	go func() { //goroutine
		<-Timers[peerTask.SubmissionID].C //waits until timer expires
		fmt.Println("Timer "+string(peerTask.SubmissionID)+" expired")
		peerTask.triggerTask() //trigger peerservice //todo extend this function
	}()

	return true
}

func stopTimer(){
	stop := Timers[2].Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
}