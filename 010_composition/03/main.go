package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CF01", "Introduction to Programming", "14"},
				course{"CF02", "Introduction to Golang", "20"},
				course{"CF03", "Introduction to Web in Go", "12"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CS01", "Mobile apps in Go", "14"},
				course{"CS02", "Advanced Golang", "20"},
				course{"CS03", "Advanced Web in Go", "12"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{"CSS01", "Go Bootcamp", "14"},
				course{"CSS02", "Go Interviews", "20"},
				course{"CSS03", "Earning the Go job", "12"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
