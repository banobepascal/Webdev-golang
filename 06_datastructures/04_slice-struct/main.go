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

	Brooke := sage{
		Name:     "Brooke",
		Position: "Cpp Developer",
	}

	Bonny := sage{
		Name:     "Bonny",
		Position: "Java Developer",
	}

	Francis := sage{
		Name:     "Francis",
		Position: "Javascript Developer",
	}

	sages := []sage{Pascal, Bonny, Brooke, Francis}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
