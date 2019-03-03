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
	const foo = `{{PRETTYTIME .}}`

	data, err := time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")
	if err != nil {
		log.Fatalf("time parsing: %s", err)
	}

	tmpl, err := template.New("deadlineDueTest").Funcs(plugin.PrettyTime()).Parse(foo)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, data)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	result := buffer.String()

	if result != "15:04 01/02/2019" {
		t.Fail()
	}
}
