package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	sage := map[string]string{
		"Pascal":  "Golang",
		"Brooke":  "Java",
		"Bonny":   "Cpp",
		"Blair":   "Fun",
		"Francis": "Javascript",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sage)
	if err != nil {
		log.Fatalln(err)
	}
}
