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

func main() {
	sage := []string{"Pascal", "Steven", "Gloria", "Paula"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sage)
	if err != nil {
		log.Fatalln(err)
	}
}
