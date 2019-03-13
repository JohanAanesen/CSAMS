package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	const parse = `{{ATOI .}}`
	const input = "42"
	const expected = 42

	tmpl, err := template.New("atoiTest").Funcs(plugin.Atoi()).Parse(parse)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, input)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result, err := strconv.Atoi(buffer.String())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if result != expected {
		t.Logf("expected: %v, got: %v", expected, result)
		t.Fail()
	}
}