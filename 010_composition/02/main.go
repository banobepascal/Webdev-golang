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

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("index.html"))

}

func main() {

	p1 := doubleZero{
		person{
			Name: "I am Bond",
			Age:  42,
		},
		true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
