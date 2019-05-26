package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	sage := map[string]string{
		"Pascal":  "Golang",
		"Brooke":  "Java",
		"Bonny":   "Cpp",
		"Blair":   "Fun",
		"Francis": "Javascript",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sage)
	if err != nil {
		log.Fatalln(err)
	}
}
