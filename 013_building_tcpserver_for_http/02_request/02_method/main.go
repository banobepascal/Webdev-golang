package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type rolex int

func (m rolex) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.html", data)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	var d rolex
	http.ListenAndServe(":8080", d)

}
