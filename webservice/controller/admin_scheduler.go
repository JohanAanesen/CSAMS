package controller

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"net/http"
	"os"
	"time"
)

func AdminSchedulerGET(w http.ResponseWriter, r *http.Request) {
	type PeerLoad struct {
		Authentication string `json:"authentication"`
		AssignmentID   int    `json:"assignment_id"`
		SubmissionID   int    `json:"submission_id"`
		Reviewers      int    `json:"reviewers"`
	}

	type PayLoad struct {
		ID             int       `json:"id"`
		Authentication string    `json:"authentication"`
		ScheduledTime  time.Time `json:"scheduled_time"`
		Task           string    `json:"task"`
		AssignmentID   int       `json:"assignment_id"`
		SubmissionID   int       `json:"submission_id"`
		Data           PeerLoad  `json:"data"`
	}

	var payloads []PayLoad

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/scheduler/index"

	resp, err := http.Get("http://" + os.Getenv("SCHEDULE_SERVICE"))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&payloads)
	if err != nil {
		panic(err)
	}

	/*
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	*/

	v.Vars["Payloads"] = payloads

	v.Render(w)
}
