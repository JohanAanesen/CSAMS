package controller

import (
	"errors"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	_ "github.com/go-sql-driver/mysql" //database driver
	"log"
	"net/http"
	"time"
)

// AdminGET handles GET-request at /admin
func AdminGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/index"

	assignmentRepo := model.AssignmentRepository{}
	assignments, err := assignmentRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	courses, err := model.GetCoursesToUser(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	v.Vars["Courses"] = courses
	v.Vars["Assignments"] = assignments

	v.Render(w)
}

// DatetimeLocalToRFC3339 converts a string from datetime-local HTML input-field to time.Time object
func DatetimeLocalToRFC3339(str string) (time.Time, error) {
	// TODO (Svein): Move this to a utils.go or something
	if str == "" {
		return time.Time{}, errors.New("error: could not parse empty datetime-string")
	}
	if len(str) < 16 {
		return time.Time{}, errors.New("cannot convert a string less then 16 characters: DatetimeLocalToRFC3339()")
	}
	year := str[0:4]
	month := str[5:7]
	day := str[8:10]
	hour := str[11:13]
	min := str[14:16]

	value := fmt.Sprintf("%s-%s-%sT%s:%s:00Z", year, month, day, hour, min)
	return time.Parse(time.RFC3339, value)
}

// GoToHTMLDatetimeLocal converts time.Time object to 'datetime-local' in HTML
func GoToHTMLDatetimeLocal(t time.Time) string {
	day := fmt.Sprintf("%d", t.Day())
	month := fmt.Sprintf("%d", t.Month())
	year := fmt.Sprintf("%d", t.Year())
	hour := fmt.Sprintf("%d", t.Hour())
	minute := fmt.Sprintf("%d", t.Minute())

	if t.Day() < 10 {
		day = "0" + day
	}

	if t.Month() < 10 {
		month = "0" + month
	}

	if t.Hour() < 10 {
		hour = "0" + hour
	}

	if t.Minute() < 10 {
		minute = "0" + minute
	}

	return fmt.Sprintf("%s-%s-%sT%s:%s", year, month, day, hour, minute)
}
