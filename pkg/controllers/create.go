package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/app/database"
	"github.com/sebastian009w/go_template/pkg/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	utils.Templates.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	conn := database.MySQL()

	if r.Method == "POST" {
		name := r.FormValue("username")
		email := r.FormValue("email")

		create, err := conn.Prepare("INSERT INTO gomi (name, email) VALUES (?, ?)")

		if err != nil {
			panic(err.Error())
		}
		create.Exec(name, email)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}
