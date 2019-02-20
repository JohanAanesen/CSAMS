package main

import (
	"log"
	"math/rand"
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

type Pairs struct{
	Items []SubPair
}

type SubPair struct{
	UserID       int `json:"userid"`
	AssignmentID int `json:"assignmentid"`
	SubmissionID1 int `json:"submissionid1"`
	SubmissionID2 int `json:"submissionid2"`
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

func createReviewPair(pairs Pairs)bool{

	tx, err := GetDB().Begin()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	for _, pair := range pairs.Items {
		_, err := tx.Exec("INSERT INTO reviewerpairs(assignmentid, userid, submissionid1, submissionid2) VALUES(?, ?, ?, ?)", pair.AssignmentID, pair.UserID, pair.SubmissionID1, pair.SubmissionID2)

		if err != nil {
			//todo log error
			log.Fatal(err.Error())
			tx.Rollback()
			return false
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func (subs Submissions) shuffle() Submissions{
	for i := range subs.Items {
		j := rand.Intn(i + 1)
		subs.Items[i], subs.Items[j] = subs.Items[j], subs.Items[i]
	}

	return subs
}