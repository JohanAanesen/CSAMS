package main

import (
	"log"
)

type Request struct {
	ID           int `json:"id"`
	AssignmentID int `json:"assignmentid"`
}

type Response struct {
	Ok bool	`json:"ok"`
}

type Submissions struct {
	Items []Submission `json:"items"`
}

type Submission struct {
	ID           int `json:"id"`
	UserID       int `json:"userid"`
	AssignmentID int `json:"assignmentid"`
}

func GetSubmissions(AssignmentID int) Submissions {

	//Create an empty courses array
	var submissions Submissions

	rows, err := GetDB().Query("SELECT id, userid, assignmentid FROM submissions WHERE submissions.assignmentid = ?", AssignmentID)
	if err != nil {
		log.Println(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return submissions
	}

	for rows.Next() {
		var id int
		var userid int
		var assignmentid int

		rows.Scan(&id, &userid, &assignmentid)

		// Add course to courses array
		submissions.Items = append(submissions.Items, Submission{
			ID:           id,
			UserID:       userid,
			AssignmentID: assignmentid,
		})
	}

	return submissions
}
