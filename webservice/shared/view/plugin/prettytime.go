package plugin

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"html/template"
	"time"
)

// PrettyTime formats time.Time formats to a prettier format for displaying in HTML
// Usage: {{PRETTYTIME .Deadline}}
func PrettyTime() template.FuncMap {
	f := make(template.FuncMap)

	f["PRETTYTIME"] = func(t time.Time) string {
		// Get correct timezone. When date is stored in the database, only the date and time is stored, not the timezone.
		timeZone := util.GetTimeInCorrectTimeZone()

		return t.Format("15:04 02/01/2006") + timeZone.Format(" -0700 MST") // Norwegian format with timezone behind
	}

	f["SORTABLE_TIME"] = func(t time.Time) string {
		// Get correct timezone. When date is stored in the database, only the date and time is stored, not the timezone.
		timeZone := util.GetTimeInCorrectTimeZone()

		return t.Format("2006/01/02 15:04") + timeZone.Format(" -0700 MST") // Norwegian format with timezone behind
	}

	return f
}
