package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/app/database"
	"github.com/sebastian009w/go_template/pkg/utils"
)

func Init(w http.ResponseWriter, r *http.Request) {
	conn := database.MySQL()

	create, err := conn.Prepare("INSERT INTO gomi (name, email) VALUES ('sebastian', 'sebastian@example.com')")

	utils.Templates.ExecuteTemplate(w, "index", nil)

	if err != nil {
		panic(err.Error())
	}
	create.Exec()

}
