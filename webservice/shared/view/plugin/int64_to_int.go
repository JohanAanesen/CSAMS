package plugin

import (
	"html/template"
)

// INT64toINT converts int64 to int, function inside templates
func INT64toINT() template.FuncMap {
	f := make(template.FuncMap)

	f["INT64_TO_INT"] = func(number int64) int {
		return int(number)
	}

	return f
}
