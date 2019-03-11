package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Task interface
type Task interface {
	Schedule()
	Trigger()
	Delete() bool
}

//PeerTask struct
type PeerTask struct {
	Authentication string `json:"authentication"`
	SubmissionID   int    `json:"submission_id"`
	AssignmentID   int    `json:"assignment_id"`
	Reviewers      int    `json:"reviewers"`
}

//Trigger runs tasks when their scheduled time expires
func (peer PeerTask) Trigger() {
	fmt.Printf("Triggering task: %v\n", peer.SubmissionID) //todo remove this

	//Send request to peer service
	jsonValue, _ := json.Marshal(peer) //json encode the request
	response, err := http.Post("http://"+os.Getenv("PEER_SERVICE"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request to peerservice failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//remove task from database
	if peer.Delete() {
		fmt.Printf("Successfully deleted task with SubmissionID: %v\n", peer.SubmissionID)
	}
}

//Schedule schedules a PeerTask for being triggered in the future
func (peer PeerTask) Schedule(scheduledTime time.Time) bool {

	payload := GetPayload(peer.SubmissionID, peer.AssignmentID)

	loc, err := time.LoadLocation("Europe/Oslo")
	if err != nil {
		log.Println("Something wrong with time location")
		return false
	}

	// TODO time
	timeNow := time.Now().In(loc)          //time now
	Duration := scheduledTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	if Duration < 0 { //scheduled time has to be in the future
		log.Printf("Could not schedule timer for submissionID: %v", peer.SubmissionID)
		peer.Delete() //todo trigger tasks that hasn't been triggered?
		return false
	}

	//afterFunc will run the function after the duration has passed
	Timers[payload.ID] = time.AfterFunc(Duration, peer.Trigger)

	log.Printf("Timer %v started with duration %v\n", payload.ID, Duration)

	return true
}

//Delete deletes a PeerTask from database
func (peer PeerTask) Delete() bool {
	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("DELETE FROM schedule_tasks WHERE submission_id = ? AND assignment_id = ?", peer.SubmissionID, peer.AssignmentID)
	if err != nil {
		//todo log error
		log.Println(err.Error())
		if err = tx.Rollback(); err != nil { //quit transaction if error
			log.Fatal(err.Error()) //die
		}
		return false
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	delete(Timers, peer.AssignmentID) //delete id from map so it may be re-assigned

	return true
}

//NewTask registers a new Task in database and ships it off for scheduling
func NewTask(payload Payload) bool {
	//Make sure a timer does not exist for this submission
	pay2 := GetPayload(payload.SubmissionID, payload.AssignmentID)
	if pay2.ID != 0{
		log.Println("Payload with these id's already exists.")
		return false
	}

	//Save the task to database for redundancy
	if !payload.Save() {
		log.Println("Something went wrong saving task")
		return false
	}

	payload = GetPayload(payload.SubmissionID, payload.AssignmentID) //get a fresh payload with id from db

	//schedule task
	if !ScheduleTask(payload) {
		log.Println("Something went wrong scheduling task")
		return false
	}

	//success
	return true
}

//ScheduleTask schedules a task based on its type
func ScheduleTask(payload Payload) bool {
	//switch based on type of task
	switch payload.Task {
	case "peer":
		peerTask, err := payload.GetPeerTask()
		if err != nil {
			log.Println("Something went wrong getting peerTask from payload")
			return false
		}

		//Schedule task
		if !peerTask.Schedule(payload.ScheduledTime) {
			log.Printf("Could not schedule task for submissionID: %v and assignmentID: %v\n", peerTask.SubmissionID, peerTask.AssignmentID)
			return false
		}
	default:
		return false
	}

	return true //success
}
