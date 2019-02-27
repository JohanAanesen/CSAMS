package model

import (
	"fmt"
	"time"
)

//Timers slice hold all current timers
var Timers  = make(map[int]*time.Timer)


func StopTimer(timerID int){
	stop := Timers[timerID].Stop()
	if stop {
		fmt.Printf("Timer %v stopped", timerID)
	}
}

func UpdateTimer(timerID int, newTime time.Time, task PeerTask) bool{
	stop := Timers[timerID].Stop()
	if stop {
		fmt.Printf("Timer %v stopped", timerID)
	}
	timeNow := time.Now() //time now
	Duration := newTime.Sub(timeNow) //subtract now's time from target time to get time until trigger

	if Duration < 0{ //scheduled time has to be in the future
		return false
	}

	Timers[task.SubmissionID] = time.AfterFunc(Duration, task.triggerTask)

	return true
}