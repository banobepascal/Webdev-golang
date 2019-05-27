package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

// To run html page
// go to terminal and type
// "go run main.go > file/name/ofchoice.html"
// file is created and then run the created file
