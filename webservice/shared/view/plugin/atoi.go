package plugin

import (
	"html/template"
	"strconv"
)

// Atoi makes an atoi function inside templates
func Atoi() template.FuncMap {
	f := make(template.FuncMap)

	f["ATOI"] = func(input string) int {
		output, _ := strconv.Atoi(input)
		return output
	}

	return f
}
