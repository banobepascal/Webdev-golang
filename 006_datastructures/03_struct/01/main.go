package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name     string
	Position string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	Pascal := sage{
		Name:     "Pascal",
		Position: "Golang Developer",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", Pascal)
	if err != nil {
		log.Fatalln(err)
	}
}
