package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"testing"
	"time"
)

func TestPrettyTime(t *testing.T) {
	const parse = `{{PRETTYTIME .}}`
	const expected = "15:04 02/01/2019"

	input, err := time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")
	if err != nil {
		log.Fatalf("time parsing: %s", err)
	}

	tmpl, err := template.New("deadlineDueTest").Funcs(plugin.PrettyTime()).Parse(parse)
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
		t.Fail()
	}
}
