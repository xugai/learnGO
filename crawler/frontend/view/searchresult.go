package view

import (
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func GetSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(fileName)),
	}
}

func (s SearchResultView) Render(out io.Writer, data interface{}) error {
	return s.template.Execute(out, data)
}