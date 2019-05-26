package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK UR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")

}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("Shida")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}

	c3, err := req.Cookie("Pascal")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}

}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Shida",
		Value: "This is Shida",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "Pascal",
		Value: "This is Pascal",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in Mozilla go to: dev tools / storage / cookies")
}
