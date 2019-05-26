package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/good.jpg", dogPic)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="good.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "good.jpg")
}
