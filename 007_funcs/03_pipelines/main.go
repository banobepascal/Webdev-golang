package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

var fm = template.FuncMap{
	"db":  double,
	"sq":  square,
	"sqr": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
