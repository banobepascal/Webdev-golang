package main

import (
	"fmt"
	"net/http"
)

type rolex int

func (m rolex) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Pascal-key", "this is from Banobe")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func </h1>")
}

func main() {

	var d rolex
	http.ListenAndServe(":8080", d)

}
