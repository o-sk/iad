package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/o-sk/gis"
)

var index = template.Must(template.ParseFiles("dist/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := index.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	i, err := gis.Search(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(j)
}

func main() {
	port := os.Args[1]
	http.HandleFunc("/api/", apiHandler)

	fs := http.FileServer(http.Dir("dist/"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, nil)
}
