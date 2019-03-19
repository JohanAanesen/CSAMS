package plugin

import (
	"html/template"
)

// Contains checks is a value is inside an array
func Contains() template.FuncMap {
	f := make(template.FuncMap)

	f["CONTAINS"] = func(e string, s []string) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}

		return false
	}

	return f
}
