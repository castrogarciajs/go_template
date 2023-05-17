package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/app/database"
	"github.com/sebastian009w/go_template/pkg/utils"
)

func Update(w http.ResponseWriter, r *http.Request) {
	conn := database.MySQL()
	ID := r.URL.Query().Get("id")

	array, err := conn.Query("SELECT * FROM gomi WHERE id = ?", ID)

	if err != nil {
		panic(err.Error())
	}

	username := utils.Username{}

	for array.Next() {
		var id int
		var name, email string

		err = array.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		username.ID = id
		username.Name = name
		username.Email = email
	}

}
