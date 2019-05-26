package main

import (
	"log"
	"os"
	"text/template"
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
		Title:   "Nothing Escaped",
		Heading: "Nothing Escaped with text/template",
		Input:   `<script>alert("NOPE");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", page)
	if err != nil {
		log.Fatalln(err)
	}

}
