package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdate": monthOfYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func monthOfYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
