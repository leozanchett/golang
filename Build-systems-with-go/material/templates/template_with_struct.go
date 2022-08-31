package main

import (
	"html/template"
	"os"
)

type User struct {
	Name   string
	UserId string
	Email  string
}

const MSG = `Dear {{.Name}},
You were registered with id {{.UserId}} 
and email {{.Email}}.`

func main() {
	u := User{
		Name:   "John Doe",
		UserId: "jdoe",
		Email:  "john@gmail.com"}

	t := template.Must(template.New("msg").Parse(MSG))
	err := t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
