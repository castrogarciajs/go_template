package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {

	http.HandleFunc("/", Init)

	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func Init(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index", nil)
}
