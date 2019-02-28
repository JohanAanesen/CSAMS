package model

import (
	"log"
	"time"
)

//Task interface
type Task interface {
	Schedule()
	TriggerTask()
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

func (peer PeerTask) Schedule(scheduledTime time.Time)bool {
	timeNow := time.Now() //time now
	Duration := scheduledTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	if Duration < 0{ //scheduled time has to be in the future
		log.Printf("Could not schedule timer for submissionID: %v", peer.SubmissionID)
		return false
	}

	//afterFunc will run the function after the duration has passed
	Timers[peer.SubmissionID] = time.AfterFunc(Duration, peer.TriggerTask)

	return true
}

func SchedulePeerTask(payload Payload)bool{
	peerTask, err := payload.GetPeerTask()
	if err != nil{
		log.Println("Something went wrong getting peerTask from payload")
		return false
	}

	//Schedule task
	if !peerTask.Schedule(payload.ScheduledTime){
		log.Printf("Could not schedule task for submissionID: %v", peerTask.SubmissionID)
		return false
	}

	//Since timer is setup nicely we want to store a reference to it and the peerTask in database for redundancy
	payload.Save(peerTask.SubmissionID)

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