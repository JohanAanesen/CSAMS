package plugin

import (
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"html/template"
)

// Atoi makes an atoi function inside templates
func GetLogType() template.FuncMap {
	f := make(template.FuncMap)

	f["GET_LOG_TYPE"] = func(activityID model.Activity) string {
		var output string
		if activityID <= model.ChangePasswordEmail && activityID >= model.NewUser {
			output = "SYSTEM"
		} else if activityID <= model.KickedFromGroup && activityID >= model.CreateSubmission {
			output = "COURSE"
		} else if activityID <= model.AdminCreateGroup && activityID >= model.AdminCreateAssignment {
			output = "ADMIN"
		} else {
			output = "N/A"
		}
		return output
	}

	return f
}
