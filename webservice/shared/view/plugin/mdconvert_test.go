package plugin_test

import (
	"bytes"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view/plugin"
	"html/template"
	"log"
	"testing"
)

func TestMDConvert(t *testing.T) {
	const parse = `{{MDCONVERT .}}`
	const input = `# Hello World`
	const expected = `<h1><a name="hello-world" class="anchor" href="#hello-world" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>Hello World</h1>
`

	tmpl, err := template.New("mdConvertTest").Funcs(plugin.MDConvert()).Parse(parse)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, input)
	if err != nil {
		log.Fatalf("execute: %s", err)
	}

	result := buffer.String()

	if result != expected {
		log.Fatalf("md convert error.\nexpected: \"%s\"\ngot: \"%s\"", expected, result)
	}
}
