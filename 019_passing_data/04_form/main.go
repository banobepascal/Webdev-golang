package main

import (
	"html/template"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {

	person := struct {
		FirstName  string
		LastName   string
		Subscribed bool
	}{
		req.FormValue("first"),
		req.FormValue("last"),
		req.FormValue("subscribe") == "on",
	}

	err := tpl.ExecuteTemplate(w, "index.html", person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
