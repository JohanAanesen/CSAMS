package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"io/ioutil"
	"net/http"
	"os"
)

func AdminSchedulerGET(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/scheduler/index"


	resp, err := http.Get("http://"+os.Getenv("SCHEDULE_SERVICE"))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	v.Vars["Content"] = string(html)

	v.Render(w)
}
