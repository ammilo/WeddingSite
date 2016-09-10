package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/bytetrip/cordoctet/models"
	"github.com/bytetrip/cordoctet/utils"
)

// RootController handles main homepage functions
func RootController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	utils.CheckErr(err)
	p := models.Page{Title: "home!"}
	t.Execute(w, p)
	fmt.Println(r.Host)
	// Simple page serve: fmt.Fprintf(w, "homepage!")
}
