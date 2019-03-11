package plugin

import (
	"html/template"
)

// Increment takes in an int and returns is plus one
func Increment() template.FuncMap {
	f := make(template.FuncMap)

	f["INCREMENT"] = func(i int) int {
		return i + 1
	}

	return f
}
