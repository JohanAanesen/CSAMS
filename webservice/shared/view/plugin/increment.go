package plugin

import (
	"html/template"
)

func Increment() template.FuncMap {
	f := make(template.FuncMap)

	f["INCREMENT"] = func(i int) int {
		return i + 1
	}

	return f
}
