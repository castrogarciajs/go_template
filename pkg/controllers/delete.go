package controllers

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	fmt.Println(ID)
}
