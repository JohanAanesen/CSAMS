package plugin

import (
	"html/template"
	"time"
)

func PrettyTime() template.FuncMap {
	f := make(template.FuncMap)

	f["PRETTYTIME"] = func(t time.Time) string {
		return t.Format("15:04 01/02/2006")
	}

	return f
}
