package main

import (
	"io"
	"net/http"
)

type rolex int

func (m rolex) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Pets are also dogs")
	case "/cat":
		io.WriteString(w, "Pets are also cats")
	}
}

func main() {

	var d rolex
	http.ListenAndServe(":8080", d)

}
