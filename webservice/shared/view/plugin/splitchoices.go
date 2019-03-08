package plugin

import (
	"html/template"
	"strings"
)

func SplitChoices() template.FuncMap {
	f := make(template.FuncMap)

	f["SPLIT_CHOICES"] = func(input string) []string {
		output := strings.Split(input, ",")

		return output
	}

	return f
}
