package main

import (
	"html/template"
	"log"
	"net/http"
)

type rolex int

func (m rolex) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", req.Form)

}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	var d rolex
	http.ListenAndServe(":8080", d)

}
