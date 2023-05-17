package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/pkg/utils"
)

func Init(w http.ResponseWriter, r *http.Request) {

	utils.Templates.ExecuteTemplate(w, "index", nil)

}
