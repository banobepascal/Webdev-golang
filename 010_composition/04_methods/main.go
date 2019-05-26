package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 9
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) AgeDivide() int {
	return p.Age / 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	p := person{
		Name: "John Wick",
		Age:  42,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
