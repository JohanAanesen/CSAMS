package model

import (
	"encoding/json"
	"github.com/JohanAanesen/CSAMS/schedulerservice/db"
	"log"
	"time"
)

//Payload struct
//https://stackoverflow.com/questions/28254102/how-to-unmarshal-json-into-interface-in-go
type Payload struct {
	ID             int             `json:"id"`
	Authentication string          `json:"authentication"`
	ScheduledTime  time.Time       `json:"scheduled_time"`
	Task           string          `json:"task"`
	AssignmentID   int             `json:"assignment_id"`
	Data           json.RawMessage `json:"data"`
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

	_, err = tx.Exec("INSERT INTO schedule_tasks(assignment_id, scheduled_time, task, data) VALUES(?, ?, ?, ?)",
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

	err = tx.Commit() //finish transaction
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}

//GetPayloads from db
func GetPayloads() []Payload {

	var payloads []Payload

	rows, err := db.GetDB().Query("SELECT id, assignment_id, scheduled_time, task, data FROM schedule_tasks")
	if err != nil {
		log.Fatal(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return []Payload{}
	}

	for rows.Next() {
		var ID int
		var assignmentID int
		var scheduledTime time.Time
		var task string
		var data []byte

		err := rows.Scan(&ID, &assignmentID, &scheduledTime, &task, &data)
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
			AssignmentID:  assignmentID,
			Data:          data,
		}

		payloads = append(payloads, payload) //add payload to slice

	}

	return payloads
}

//GetPayload from db
func GetPayload(assID int) Payload {

	rows, err := db.GetDB().Query("SELECT id, assignment_id, scheduled_time, task, data FROM schedule_tasks WHERE assignment_id = ?", assID)
	if err != nil {
		log.Fatal(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return Payload{}
	}

	for rows.Next() {
		var ID int
		var assignmentID int
		var scheduledTime time.Time
		var task string
		var data []byte

		err := rows.Scan(&ID, &assignmentID, &scheduledTime, &task, &data)
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
			AssignmentID:  assignmentID,
			Data:          data,
		}

		return payload //return first and best

	}

	return Payload{}
}

//DeletePayload removes payload from db
func DeletePayload(assID int) bool {

	payload := GetPayload(assID)

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("DELETE FROM schedule_tasks WHERE assignment_id = ?", assID)

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

	StopTimer(payload.ID) //stops ongoing timer

	return true
}

//UpdatePayload updates payload in db
func (payload Payload) UpdatePayload() bool {

	tx, err := db.GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = tx.Exec("UPDATE schedule_tasks SET scheduled_time = ?, task = ?, data = ? WHERE assignment_id = ?",
		payload.ScheduledTime,
		payload.Task,
		payload.Data,
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
