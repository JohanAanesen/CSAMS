package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"html/template"
	"log"
	"testing"
)

func TestSplitChoices(t *testing.T) {
	const parse = `{{SPLIT_CHOICES .}}`
	const input = "a,b,c"
	const expected = "[a b c]"

	tmpl, err := template.New("splitChoicesTest").Funcs(plugin.SplitChoices()).Parse(parse)
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
