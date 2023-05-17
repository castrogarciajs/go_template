package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/app/database"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	conn := database.MySQL()
	delete, err := conn.Prepare("DELETE FROM gomi WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(ID)
	http.Redirect(w, r, "/", http.StatusNoContent)
}
