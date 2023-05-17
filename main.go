package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Init)

	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func Init(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
