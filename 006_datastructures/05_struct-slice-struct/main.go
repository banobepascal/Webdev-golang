package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name     string
	Position string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	P := sage{
		Name:     "Pascal",
		Position: "Golang Developer",
	}

	Bro := sage{
		Name:     "Brooke",
		Position: "Cpp Developer",
	}

	Bon := sage{
		Name:     "Bonny",
		Position: "Java Developer",
	}

	Fra := sage{
		Name:     "Francis",
		Position: "Javascript Developer",
	}

	one := car{
		Manufacturer: "Toyota",
		Model:        "Rav4",
		Doors:        2,
	}

	two := car{
		Manufacturer: "Ford",
		Model:        "Mustang",
		Doors:        2,
	}

	three := car{
		Manufacturer: "VolksWagen",
		Model:        "Golf",
		Doors:        4,
	}

	sages := []sage{P, Bon, Bro, Fra}
	cars := []car{one, two, three}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
