package plugin

import (
	"fmt"
	"html/template"
)

// Sprintf func
func Sprintf() template.FuncMap {
	f := make(template.FuncMap)

	f["Sprintf"] = func(format string, args ...interface{}) string {
		return fmt.Sprintf(format, args...)
	}

	return f
}
