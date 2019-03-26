package controller

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/scheduler"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//AdminSchedulerGET handles GET requests to /admin/scheduler
func AdminSchedulerGET(w http.ResponseWriter, r *http.Request) {
	type PeerLoad struct {
		Authentication string `json:"authentication"`
		AssignmentID   int    `json:"assignment_id"`
		Reviewers      int    `json:"reviewers"`
	}

	type PayLoad struct {
		ID             int       `json:"id"`
		Authentication string    `json:"authentication"`
		ScheduledTime  time.Time `json:"scheduled_time"`
		Task           string    `json:"task"`
		AssignmentID   int       `json:"assignment_id"`
		Data           PeerLoad  `json:"data"`
	}

	var payloads []PayLoad

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/scheduler/index"

	url := "http://localhost:8086" //schedulerservice

	if os.Getenv("SCHEDULE_SERVICE") != "" {
		url = "http://" + os.Getenv("SCHEDULE_SERVICE") //schedulerservice address changed in env var
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&payloads)
	if err != nil {
		panic(err)
	}

	v.Vars["Payloads"] = payloads

	v.Render(w)
}

//AdminSchedulerDELETE handles POST requests to the delete address
func AdminSchedulerDELETE(w http.ResponseWriter, r *http.Request) {
	assIDString := r.FormValue("assid")

	if assIDString == ""{
		log.Println("Either assid or subid was not provided")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	assID, err := strconv.Atoi(assIDString)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}


	sched := scheduler.Scheduler{}

	if sched.SchedulerExists(assID) {
		err := sched.DeleteSchedule(assID)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/admin/scheduler", http.StatusFound)
}
