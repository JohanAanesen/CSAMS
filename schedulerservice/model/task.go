package model

import (
	"log"
	"time"
)

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

func (peer PeerTask) TriggerTask() {
	//todo send request to peerservice
}

func SchedulePeerTask(payload Payload)bool{
	peerTask, err := payload.getPeerTask()
	if err != nil{
		log.Println("Something went wrong getting peerTask from payload")
		return false
	}

	timeNow := time.Now() //time now
	Duration := payload.ScheduledTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	if Duration < 0{ //scheduled time has to be in the future
		return false
	}

	//afterFunc will run the function after the duration has passed
	Timers[peerTask.SubmissionID] = time.AfterFunc(Duration, peerTask.TriggerTask)

	/*
		Timers[peerTask.SubmissionID] = time.AfterFunc(time.Second*10, func() {
			fmt.Println("Task running yei")
		})


		Timers[peerTask.SubmissionID] = time.NewTimer(Duration) //register timer
		go func() { //goroutine aka lightweight 'thread'
			<-Timers[peerTask.SubmissionID].C //waits until timer expires
			fmt.Printf("Timer %v expired", peerTask.SubmissionID)
			peerTask.triggerTask() //trigger peerservice //todo extend this function
		}()
	*/
	return true //success
}