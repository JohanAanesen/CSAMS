package plugin

import (
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
)

func MDConvert() template.FuncMap {
	f := make(template.FuncMap)

	f["MDCONVERT"] = func(input string) template.HTML {
		md := []byte(input)
		rendered := github_flavored_markdown.Markdown(md)
		return template.HTML(rendered)
	}

	return f
}
