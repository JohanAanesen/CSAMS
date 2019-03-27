package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"regexp"
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

	// TODO time-norwegian +0100 CET or +0200 CEST
	expected := regexp.MustCompile("^15:04 02/01/2019 &#43;0[1|2]00 CE[|S]?T$")
	if !expected.Match([]byte(result)) {
		t.Errorf("\nexpected:\t%v\ngot:\t\t%v", expected.String(), result)
		t.Fail()
	}
}
