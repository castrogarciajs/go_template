package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/pkg/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	utils.Templates.ExecuteTemplate(w, "Create", nil)
}
