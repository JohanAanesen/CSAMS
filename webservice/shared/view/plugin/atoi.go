package plugin

import (
	"html/template"
	"strconv"
)

//Atoi function
func Atoi() template.FuncMap {
	f := make(template.FuncMap)

	f["ATOI"] = func(input string) int {
		output, _ := strconv.Atoi(input)
		return output
	}

	return f
}
