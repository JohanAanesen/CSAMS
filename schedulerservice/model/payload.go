package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/db"
	"encoding/json"
	"log"
	"time"
)

//Payload struct
//https://stackoverflow.com/questions/28254102/how-to-unmarshal-json-into-interface-in-go
type Payload struct {
	ScheduledTime time.Time `json:"scheduled_time"`
	Task          string    `json:"task"`
	SubmissionID  int       `json:"submission_id"`
	Data          json.RawMessage
}

func (payload Payload) GetPeerTask() (PeerTask, error) {
	var peerTask PeerTask
	if err := json.Unmarshal([]byte(payload.Data), &peerTask); err != nil { //read the unread json into peerTask
		return PeerTask{}, err
	}

	return peerTask, nil
}

func (payload Payload) Save() bool {
	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("INSERT INTO schedule_tasks(submission_id, scheduled_time, task, data) VALUES(?, ?, ?, ?)",
		payload.SubmissionID,
		payload.ScheduledTime,
		payload.Task,
		payload.Data)

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

	return true
}
