package model

import (
	"os"
	"text/template"
)

type Form struct {
	Action  string
	Method  string
	EncType string
	Datum   []Input
}

func (f Form) Generate() error {
	tempBody := `<form action="{{.Action}}"  method="{{.Method}}" enctype="{{.EncType}}">
{{range $_,$tipe := .Datum}}  <input name="{{.Name}}" type="{{.Type}}">
{{end}}  <input type="submit">
</form>`

	temp := template.Must(template.New("").Parse(tempBody))
	return temp.Execute(os.Stdout, f)
}

type Input struct {
	Name string
	Type string
}

func (i *Input) New(name, types string) {
	i.Name = name
	i.Type = types
}
