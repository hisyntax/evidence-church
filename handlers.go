package main

import (
	"net/http"
	"text/template"
)

var index = template.Must(template.ParseFiles("templates/index.html"))
var sermon = template.Must(template.ParseFiles("templates/sermon.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func sermonHandler(w http.ResponseWriter, r *http.Request) {
	sermon.Execute(w, nil)
}
