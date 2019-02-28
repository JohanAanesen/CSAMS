package model

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/db"
	"log"
	"time"
)

//Timers slice hold all current timers
var Timers  = make(map[int]*time.Timer)

func GetTimer(timerID int)*time.Timer{
	return Timers[timerID]
}

func StopTimer(timerID int){
	stop := Timers[timerID].Stop()
	if stop {
		fmt.Printf("Timer %v stopped\n", timerID)
	}
}

func UpdateTimer(timerID int, newTime time.Time, task PeerTask) bool{
	stop := Timers[timerID].Stop()
	if stop {
		fmt.Printf("Timer %v stopped\n", timerID)
	}
	timeNow := time.Now() //time now
	Duration := newTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	if Duration < 0{ //scheduled time has to be in the future
		return false
	}

	Timers[task.SubmissionID] = time.AfterFunc(Duration, task.Trigger)

	//todo update db

	return true
}

func InitializeTimers(){
	rows, err := db.GetDB().Query("SELECT submission_id, scheduled_time, task, data FROM schedule_tasks")
	if err != nil {
		log.Fatal(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return
	}

	for rows.Next() {
		var submissionID int
		var scheduledTime time.Time
		var task string
		var data []byte

		err := rows.Scan(&submissionID, &scheduledTime, &task, &data)
		if err != nil {
			log.Println(err.Error()) // TODO : log error
			// returns empty course array if it fails
			return
		}

		// Add course to courses array
		var payload Payload

		payload = Payload{
			ScheduledTime:scheduledTime,
			Task:task,
			SubmissionID:submissionID,
			Data:data,
		}

		if payload.ScheduledTime.Sub(time.Now()) < 0{ //trigger tasks that has dinged when service was down
			task, err := payload.GetPeerTask()
			if err != nil{
				log.Printf("asdla") //todo
			}

			task.Trigger()
			return
		} else if !ScheduleTask(payload){ //schedule task
			log.Printf("Could not initialize timer for submission ID: %v\n", submissionID)
		}
	}
}