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
		Method        string
		URL           *url.URL
		Submissions   url.Values
		ContentLength int64
		Host          string
		Header        http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.ContentLength,
		req.Host,
		req.Header,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	var d rolex
	http.ListenAndServe(":8080", d)

}
