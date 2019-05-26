package main

import (
	"html/template"
	"log"
	"os"
)

type pages struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("index.html"))

}

func main() {

	page := pages{
		Title:   "Escaped",
		Heading: "Danger is Escaped with html/template",
		Input:   `<script>alert("Yep!")</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", page)
	if err != nil {
		log.Fatalln(err)
	}

}
