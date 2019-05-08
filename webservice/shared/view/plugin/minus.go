package plugin

import (
	"database/sql"
	"html/template"
)

// Minus function to substract two ints
func Minus() template.FuncMap {
	f := make(template.FuncMap)

	f["MINUS"] = func(a int, b int) int {
		return a - b
	}

	return f
}

// NullInt64MinusInt function to substract two ints
func NullInt64MinusInt() template.FuncMap {
	f := make(template.FuncMap)

	f["NULLINT64MINUSINT"] = func(a sql.NullInt64, b int) int {
		return int(a.Int64) - b
	}

	return f
}
