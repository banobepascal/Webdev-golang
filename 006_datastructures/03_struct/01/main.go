package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name     string
	Position string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	Pascal := sage{
		Name:     "Pascal",
		Position: "Golang Developer",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", Pascal)
	if err != nil {
		log.Fatalln(err)
	}
}
