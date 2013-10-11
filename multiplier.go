package main

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	//"regexp"
)

type matrixArrangement struct {
	ARows int
	AColumns int
	BColumns int
}

/*
func arrangeMatrices (arows int, acolumns int, bcolumns int) (*matrixArrangement) {
	m := &matrixArrangement{ARows: arows, AColumns: acolumns, BColumns: bcolumns};
	return m;
}
*/

var templates = template.Must (template.ParseFiles ("index.html"))

func renderTemplate (w http.ResponseWriter, tmpl string, m *matrixArrangement) {
	err := templates.ExecuteTemplate (w, tmpl + ".html", m)
    if err != nil {
    	http.Error (w, err.Error (), http.StatusInternalServerError)
        return
    }
}

func valueConvert(val string) (int, error) {
	value, err := strconv.Atoi (val);
	if err != nil {
		fmt.Println("A conversion error occurred: %s", err);
		return 0, err
	}

	return value, nil
}

func (m *matrixArrangement) viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("View");

	init := "index";
	
	renderTemplate (w, init, m);
}

func (m *matrixArrangement) defineHandler(w http.ResponseWriter, r *http.Request) {
	a, err1 := valueConvert(r.FormValue("aRows"));
	
	if err1 != nil {
		http.Redirect(w, r, "/", http.StatusFound);
		fmt.Println("Please evaluate input for errors and resubmit");
		return
	}
	
	b, err2 := valueConvert(r.FormValue("aColumns"));

	if err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound);
		fmt.Println("Please evaluate input for errors and resubmit");
		return
	}

	c, err3 := valueConvert(r.FormValue("bColumns"));

	if err3 != nil {
		http.Redirect(w, r, "/", http.StatusFound);
		fmt.Println("Please evaluate input for errors and resubmit");
		return
	}

	m.ARows = a;
	m.AColumns = b;
	m.BColumns = c;

	m.viewHandler(w, r)
}

func main() {
	Init := &matrixArrangement {ARows: 0, AColumns: 0, BColumns: 0}
	http.HandleFunc("/", Init.viewHandler)
	http.HandleFunc("/define", Init.defineHandler)
	http.ListenAndServe(":8080", nil)
}