package plugin

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"html/template"
	"log"
	"os"
	"time"
)

// DeadlineDue returns a boolean if the time given is passed current time
// Usage: {{if DEADLINEDUE .SomeTime}}...{{end}}
func DeadlineDue() template.FuncMap {
	f := make(template.FuncMap)

	f["DEADLINEDUE"] = func(t time.Time) bool {
		// TODO time-norwegian
		loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
		if err != nil {
			log.Println(err.Error())
		}

		// TODO fix hack
		t = t.In(loc).Add(-time.Hour)
		return t.Before(util.GetTimeInCorrectTimeZone())
	}

	return f
}
