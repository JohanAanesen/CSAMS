package plugin

import (
	"html/template"
	"time"
)

// DeadlineDue returns a boolean if the time given is passed current time
// Usage: {{if DEADLINEDUE .SomeTime}}...{{end}}
func DeadlineDue() template.FuncMap {
	f := make(template.FuncMap)

	f["DEADLINEDUE"] = func(t time.Time) bool {
		// TODO time
		return t.Before(time.Now().UTC().Add(time.Hour))
	}

	return f
}
