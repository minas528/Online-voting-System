package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("ui/templates/*"))

func main() {
	fs := http.FileServer(http.Dir("ui/assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets", fs))
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8989", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}
