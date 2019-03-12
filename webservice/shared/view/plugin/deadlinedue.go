package plugin

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"html/template"
	"time"
)

// DeadlineDue returns a boolean if the time given is passed current time
// Usage: {{if DEADLINEDUE .SomeTime}}...{{end}}
func DeadlineDue() template.FuncMap {
	f := make(template.FuncMap)

	f["DEADLINEDUE"] = func(t time.Time) bool {
		// TODO time-norwegian
		return t.Before(util.GetTimeInNorwegian())
	}

	return f
}
