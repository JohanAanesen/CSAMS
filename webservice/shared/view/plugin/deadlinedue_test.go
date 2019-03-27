package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"testing"
	"time"
)

func TestDeadlineDue(t *testing.T) {
	// TODO time
	const parse = `{{DEADLINEDUE .}}`
	var input = time.Now().UTC().Add(-time.Hour)
	var expected = "true"

	tmpl, err := template.New("deadlineDueTest").Funcs(plugin.DeadlineDue()).Parse(parse)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, input)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result := buffer.String()

	if result != expected {
		t.Logf("expected: %v, got: %v", expected, result)
		t.Fail()
	}

	// TODO time
	input = time.Now().UTC().Add(+2 * time.Hour)
	expected = "false"

	buffer = new(bytes.Buffer)

	err = tmpl.Execute(buffer, input)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result = buffer.String()
	if result != expected {
		t.Logf("expected: %v, got: %v", expected, result)
		t.Fail()
	}
}
