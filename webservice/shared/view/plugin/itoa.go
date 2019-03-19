package plugin

import (
	"html/template"
	"strconv"
)

// Itoa converts an integer to a string
func Itoa() template.FuncMap {
	f := make(template.FuncMap)

	f["ITOA"] = func(input int) string {
		return strconv.Itoa(input)
	}

	return f
}
