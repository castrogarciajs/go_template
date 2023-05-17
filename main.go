package main

import (
	"log"
	"net/http"

	"github.com/sebastian009w/go_template/pkg/controllers"
)

func main() {

	http.HandleFunc("/", controllers.Init)

	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
