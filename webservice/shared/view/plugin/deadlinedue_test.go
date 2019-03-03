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
	then := time.Now().UTC().Add(-time.Hour)

	const foo = `{{DEADLINEDUE .}}`

	tmpl, err := template.New("deadlineDueTest").Funcs(plugin.DeadlineDue()).Parse(foo)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, then)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result := buffer.String()

	if result != "true" {
		t.Fail()
	}

	then = time.Now().UTC().Add(+2 * time.Hour)

	buffer = new(bytes.Buffer)

	err = tmpl.Execute(buffer, then)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result = buffer.String()
	if result != "false" {
		t.Fail()
	}
}
