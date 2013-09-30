package main

import (
	"fmt"
	"net/http"
	"html/template"
	//"regexp"
)

type matrix struct {
	id int
	Rows int
	Columns int
}

func renderTemplate (w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate (w, "templates" + tmpl + ".html", p)
    if err != nil {
    	http.Error (w, err.Error (), http.StatusInternalServerError)
        return
    }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}