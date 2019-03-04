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
	ID             int       `json:"id"`
	Authentication string    `json:"authentication"`
	ScheduledTime  time.Time `json:"scheduled_time"`
	Task           string    `json:"task"`
	AssignmentID   int       `json:"assignment_id"`
	SubmissionID   int       `json:"submission_id"`
	Data           json.RawMessage
}

//GetPeerTask withdraws a PeerTask object from the data column of a payload object
func (payload Payload) GetPeerTask() (PeerTask, error) {
	var peerTask PeerTask
	if err := json.Unmarshal([]byte(payload.Data), &peerTask); err != nil { //read the unread json into peerTask
		return PeerTask{}, err
	}

	return peerTask, nil
}

//Save stores the payload object to the database for redundancy
func (payload Payload) Save() bool {
	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	ex, err := tx.Exec("INSERT INTO schedule_tasks(submission_id, assignment_id, scheduled_time, task, data) VALUES(?, ?, ?, ?, ?)",
		payload.SubmissionID,
		payload.AssignmentID,
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

	schedId, err := ex.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	payload.ID = int(schedId)

	return true
}

//GetPayloads from db
func GetPayloads() []Payload {

	var payloads []Payload

	rows, err := db.GetDB().Query("SELECT id, submission_id, assignment_id, scheduled_time, task, data FROM schedule_tasks")
	if err != nil {
		log.Fatal(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return []Payload{}
	}

	for rows.Next() {
		var ID int
		var submissionID int
		var assignmentID int
		var scheduledTime time.Time
		var task string
		var data []byte

		err := rows.Scan(&ID, &submissionID, &assignmentID, &scheduledTime, &task, &data)
		if err != nil {
			log.Println(err.Error()) // TODO : log error
			// returns empty course array if it fails
			return []Payload{}
		}

		// Add course to courses array
		var payload Payload

		payload = Payload{
			ID:            ID,
			ScheduledTime: scheduledTime,
			Task:          task,
			SubmissionID:  submissionID,
			AssignmentID:  assignmentID,
			Data:          data,
		}

		payloads = append(payloads, payload) //add payload to slice

	}

	return payloads
}

//GetPayload from db
func GetPayload(subID int, assID int) Payload {

	rows, err := db.GetDB().Query("SELECT id, submission_id, assignment_id, scheduled_time, task, data FROM schedule_tasks WHERE submission_id = ? AND assignment_id = ?", subID, assID)
	if err != nil {
		log.Fatal(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return Payload{}
	}

	for rows.Next() {
		var ID int
		var submissionID int
		var assignmentID int
		var scheduledTime time.Time
		var task string
		var data []byte

		err := rows.Scan(&ID, &submissionID, &assignmentID, &scheduledTime, &task, &data)
		if err != nil {
			log.Println(err.Error()) // TODO : log error
			// returns empty course array if it fails
			return Payload{}
		}

		// Add course to courses array
		var payload Payload

		payload = Payload{
			ID:            ID,
			ScheduledTime: scheduledTime,
			Task:          task,
			SubmissionID:  submissionID,
			AssignmentID:  assignmentID,
			Data:          data,
		}

		return payload //return first and best

	}

	return Payload{}
}

//DeletePayload removes payload from db
func DeletePayload(subID int, assID int) bool {

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("DELETE FROM schedule_tasks WHERE submission_id = ? AND assignment_id = ?", subID, assID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
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

	StopTimer(subID) //stops ongoing timer

	return true
}

//UpdatePayload updates payload in db
func (payload Payload) UpdatePayload() bool {

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("UPDATE schedule_tasks SET scheduled_time = ?, task = ?, data = ? WHERE submission_id = ? AND assignment_id = ?",
		payload.ScheduledTime,
		payload.Task,
		payload.Data,
		payload.SubmissionID,
		payload.AssignmentID)

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
