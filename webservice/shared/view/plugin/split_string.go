package plugin

import (
	"html/template"
	"strings"
)

// SplitString splits a string into an array
func SplitString() template.FuncMap {
	f := make(template.FuncMap)

	f["SPLIT_STRING"] = func(s string, sep string) []string {
		return strings.Split(s, sep)
	}

	return f
}
