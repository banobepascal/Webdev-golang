package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ut-8")

	io.WriteString(w, `
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)

}

func main() {

	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)

}
