package model

import (
	"fmt"
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
	fmt.Printf("DING: %v", peer.SubmissionID) //todo remove this
}

func (peer PeerTask) Schedule(scheduledTime time.Time)bool {

	loc, err := time.LoadLocation("Europe/Oslo")
	if err != nil{
		log.Println("Something wrong with time location")
		return false
	}

	timeNow := time.Now().In(loc) //time now
	Duration := scheduledTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	fmt.Printf("Duration registered: %v\n", Duration) //todo remove this

	if Duration < 0{ //scheduled time has to be in the future
		log.Printf("Could not schedule timer for submissionID: %v", peer.SubmissionID)
		return false
	}

	//afterFunc will run the function after the duration has passed
	Timers[peer.SubmissionID] = time.AfterFunc(Duration, peer.TriggerTask)

	return true
}

func NewTask(payload Payload)bool{
	SubID := ScheduleTask(payload) //schedule task
	if SubID == 0{
		log.Println("Something went wrong scheduling task")
		return false
	}

	//Save the scheduled task to database for redundancy
	return SubID != 0 && payload.Save(SubID)
}

func ScheduleTask(payload Payload)int{

	var subID = 0

	switch payload.Task{
	case "peer":
		peerTask, err := payload.GetPeerTask()
		if err != nil{
			log.Println("Something went wrong getting peerTask from payload")
			return 0
		}

		//Schedule task
		if !peerTask.Schedule(payload.ScheduledTime){
			log.Printf("Could not schedule task for submissionID: %v", peerTask.SubmissionID)
			return 0
		}

		subID = peerTask.SubmissionID
	default:
		return 0
	}

	return subID //success
}