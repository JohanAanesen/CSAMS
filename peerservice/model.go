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
	UserID       int `json:"user_id"`
	SubmissionID int `json:"submission_id"`
	AssignmentID int `json:"assignment_id"`
}

//Pairs type subpair slice
type Pairs []SubPair

//SubPair struct
type SubPair struct {
	SubmissionID int `json:"submission_id"`
	AssignmentID int `json:"assignment_id"`
	UserID       int `json:"user_id"`
	ReviewUserID int `json:"review_user_id"`
}

//getSubmissions fetches user_submissions from database with submissionID
func getSubmissions(SubmissionID int, AssignmentID int) Submissions {

	//Create an empty courses array
	var submissions Submissions

	rows, err := GetDB().Query("SELECT user_id FROM user_submissions WHERE submission_id = ? AND assignment_id = ? GROUP BY user_id", SubmissionID, AssignmentID)
	if err != nil {
		log.Println(err.Error()) // TODO : log error
		// returns empty course array if it fails
		return submissions
	}

	for rows.Next() {
		var userid int

		rows.Scan(&userid)

		// Add course to courses array
		submissions = append(submissions, Submission{
			UserID:       userid,
			SubmissionID: SubmissionID,
			AssignmentID: AssignmentID,
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

		_, err := tx.Exec("INSERT INTO peer_reviews(submission_id, assignment_id, user_id, review_user_id) VALUES(?, ?, ?, ?)", pair.SubmissionID, pair.AssignmentID, pair.UserID, pair.ReviewUserID)

		if err != nil {
			//todo log error

			log.Fatal(err.Error())
			tx.Rollback() //quit transaction if error
			return false
		}

		log.Printf("Pair generated: %v", pair)
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
