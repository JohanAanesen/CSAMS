package model

import (
	"encoding/json"
	"time"
)

//Payload struct
//https://stackoverflow.com/questions/28254102/how-to-unmarshal-json-into-interface-in-go
type Payload struct {
	ScheduledTime time.Time       `json:"scheduled_time"`
	Task          string          `json:"task"`
	Data          json.RawMessage
}

func (payload Payload)GetPeerTask()(PeerTask, error){
	var peerTask PeerTask
	if err := json.Unmarshal([]byte(payload.Data), &peerTask); err != nil { //read the unread json into peerTask
		return PeerTask{}, err
	}

	return peerTask, nil
}