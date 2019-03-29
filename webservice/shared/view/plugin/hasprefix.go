package plugin

import (
	"html/template"
	"strings"
)

// HasPrefix func
func HasPrefix() template.FuncMap {
	f := make(template.FuncMap)

	f["HasPrefix"] = func(s, prefix string) bool {
		return strings.HasPrefix(s, prefix)
	}

	return f
}
