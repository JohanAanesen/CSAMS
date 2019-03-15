package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"strconv"
	"testing"
)

func TestIncrement(t *testing.T) {
	const parse = `{{INCREMENT .}}`
	const input = 3
	const expected = 4

	tmpl, err := template.New("incrementTest").Funcs(plugin.Increment()).Parse(parse)
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
		t.Fail()
	}
}
