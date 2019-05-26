package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/good.jpg", dogPic)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="good.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("good.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
