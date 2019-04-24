package plugin

import (
	"html/template"
)

// Append appends one string to the other, function inside templates
func Append() template.FuncMap {
	f := make(template.FuncMap)

	f["APPEND"] = func(input string, appeninput string) string {
		return input + appeninput
	}

	return f
}
