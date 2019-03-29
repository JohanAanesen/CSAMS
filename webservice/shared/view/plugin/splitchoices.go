package plugin

import (
	"html/template"
	"strings"
)

// SplitChoices splits a string by "," (comma) and returns a string-slice
func SplitChoices() template.FuncMap {
	f := make(template.FuncMap)

	f["SPLIT_CHOICES"] = func(input string) []string {
		output := strings.Split(input, "|")

		return output
	}

	f["Join"] = func(a []string, sep string) string {
		return strings.Join(a, sep)
	}

	return f
}
