package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	p1 := person{
		Name: "Banobe Pascal",
		Age:  21,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
