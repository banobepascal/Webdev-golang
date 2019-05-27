package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name     string
	Position string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	P1 := person{
		Name:     "Pascal",
		Position: "Golang Developer",
	}

	Q1 := person{
		Name:     "Francis",
		Position: "Javascript Developer",
	}

	R1 := person{
		Name:     "Gloria",
		Position: "Erlang Developer",
	}

	sage := []person{P1, Q1, R1}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sage)
	if err != nil {
		log.Fatalln(err)
	}
}
