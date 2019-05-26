package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func m(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("err parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.html", "Banobe Pascal")
	if err != nil {
		log.Fatalln("error somnething is wrong", err)
	}
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dogs are pretty")
}

func main() {

	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(m))

	http.ListenAndServe(":8080", nil)

}
