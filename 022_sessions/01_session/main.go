package main

import (
	"fmt"
	"html/template"
	"net/http"

	guuid "github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := guuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			// secure true
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		fmt.Printf("this is my unique id: %s\n", sID.String())
	}

	// if the user already exists
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}
