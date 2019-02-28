package plugin

import (
	"html/template"
	"time"
)

func DeadlineDue() template.FuncMap {
	f := make(template.FuncMap)

	f["DEADLINEDUE"] = func(t time.Time) bool {
		return t.Before(time.Now().UTC().Add(time.Hour))
	}

	return f
}
