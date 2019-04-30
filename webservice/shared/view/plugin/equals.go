package plugin

import (
	"database/sql"
	"html/template"
)

// Equals function similar to eq
func Equals() template.FuncMap {
	f := make(template.FuncMap)

	f["EQUALS"] = func(a int, b int) bool {
		return a == b
	}

	return f
}

// NullInt64EqualsInt function similar to eq
func NullInt64EqualsInt() template.FuncMap {
	f := make(template.FuncMap)

	f["NULLINT64EQUALSINT"] = func(a sql.NullInt64, b int) bool {
		return int(a.Int64) == b
	}

	return f
}
