package plugin

import (
	"html/template"
	"strings"
)

// SplitChoices splits a string by "," (comma) and returns a string-slice
func SplitChoices() template.FuncMap {
	f := make(template.FuncMap)

	f["SPLIT_CHOICES"] = func(input string) []string {
		output := strings.Split(input, ",")

		return output
	}

	return f
}
