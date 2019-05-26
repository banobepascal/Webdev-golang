package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 201)
	if err != nil {
		log.Fatalln(err)
	}

}
