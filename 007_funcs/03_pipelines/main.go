package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"db":  double,
	"sq":  square,
	"sqr": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
