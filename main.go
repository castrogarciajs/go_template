package main

import (
	"log"
	"net/http"

	"github.com/sebastian009w/go_template/pkg/controllers"
)

// vscode globals --commits --locals --no-locals
func main() {

	http.HandleFunc("/", controllers.Init)
	http.HandleFunc("/create", controllers.Create)

	// URL method POST
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)

	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
