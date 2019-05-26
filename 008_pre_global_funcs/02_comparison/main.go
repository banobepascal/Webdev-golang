package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type g1 struct {
	Score1 int
	Score2 int
}

func main() {

	goose := g1{
		Score1: 12,
		Score2: 13,
	}

	err := tpl.Execute(os.Stdout, goose)
	if err != nil {
		log.Fatalln(err)
	}

}
