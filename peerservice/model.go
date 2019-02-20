package main

import (
	"log"
	"math/rand"
	"time"
)

type Request struct {
	ID           int `json:"id"`
	SubmissionID int `json:"submissionid"`
	Reviewers    int `json:"reviewers"`
}

type Response struct {
	Ok bool `json:"ok"`
}

type Submissions struct {
	Items []Submission `json:"items"`
}

type Submission struct {
	ID           int `json:"id"`
	UserID       int `json:"userid"`
	SubmissionID int `json:"submissionid"`
}

type Pairs struct {
	Items []SubPair
}

type SubPair struct {
	SubmissionID int `json:"submissionid"`
	UserID       int `json:"userid"`
	ReviewID     int `json:"reviewid"`
}

func GetSubmissions(SubmissionID int) Submissions {

	//Create an empty courses array
	var submissions Submissions

	rows, err := GetDB().Query("SELECT id, user_id, submission_id FROM user_submissions WHERE submission_id = ?", SubmissionID)
	if err != nil {
		log.Println(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return submissions
	}

	for rows.Next() {
		var id int
		var userid int
		var submissionID int

		rows.Scan(&id, &userid, &submissionID)

		// Add course to courses array
		submissions.Items = append(submissions.Items, Submission{
			ID:           id,
			UserID:       userid,
			SubmissionID: submissionID,
		})
	}

	return submissions
}

func createReviewPair(pairs Pairs) bool {

	tx, err := GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	for _, pair := range pairs.Items {
		_, err := tx.Exec("INSERT INTO reviewpairs(submission_id, user_id, review_submission_id) VALUES(?, ?, ?)", pair.SubmissionID, pair.UserID, pair.ReviewID)

		if err != nil {
			//todo log error
			log.Fatal(err.Error())
			tx.Rollback() //quit transaction if error
			return false
		}
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}

func (subs Submissions) shuffle() Submissions {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range subs.Items {
		j := r.Intn(i + 1)
		subs.Items[i], subs.Items[j] = subs.Items[j], subs.Items[i]
	}

	return subs
}
