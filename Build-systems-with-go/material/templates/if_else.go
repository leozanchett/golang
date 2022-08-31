package main

import (
	"os"
	"text/template"
)

type User struct {
	Name   string
	Female bool
}

const MSGS = `
{{if .Female}}Mrs.{{else}}Mr.{{end}}{{.Name}},
Your package is ready.
Thanks
`

func main() {
	u1 := User{"John", false}
	u2 := User{"Mary", true}

	t := template.Must(template.New("msg").Parse(MSGS))
	err := t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, u2)
	if err != nil {
		panic(err)
	}
}
