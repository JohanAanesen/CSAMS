package main

import (
	"log"
	"math/rand"
	"time"
)

//Submissions type submission slice
type Submissions []Submission

//Submission struct
type Submission struct {
	ID           int `json:"id"`
	UserID       int `json:"userid"`
	SubmissionID int `json:"submissionid"`
}

//Pairs type subpair slice
type Pairs []SubPair

//SubPair struct
type SubPair struct {
	SubmissionID int `json:"submissionid"`
	UserID       int `json:"userid"`
	ReviewID     int `json:"reviewid"`
}

//getSubmissions fetches user_submissions from database with submissionID
func getSubmissions(SubmissionID int) Submissions {

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
		submissions = append(submissions, Submission{
			ID:           id,
			UserID:       userid,
			SubmissionID: submissionID,
		})
	}

	return submissions
}

//savePairs saves the peer_reviews to database
func savePairs(pairs Pairs) bool {

	tx, err := GetDB().Begin() //start transaction
	if err != nil {
		log.Println(err.Error())
		return false
	}

	for _, pair := range pairs {
		_, err := tx.Exec("INSERT INTO peer_reviews(submission_id, user_id, review_submission_id) VALUES(?, ?, ?)", pair.SubmissionID, pair.UserID, pair.ReviewID)

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

//shuffle randomly scrambles a slice
func (subs Submissions) shuffle() Submissions {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range subs {
		j := r.Intn(i + 1)
		subs[i], subs[j] = subs[j], subs[i]
	}

	return subs
}
