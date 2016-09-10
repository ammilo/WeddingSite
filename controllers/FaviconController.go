package controllers

import (
	"fmt"
	"net/http"
)

// FaviconController handles main homepage functions
func FaviconController(w http.ResponseWriter, r *http.Request) {
	// Debugging helper
	fmt.Println("favicon.ico")

	// Right now it's just a redirect to /static/img/favicon.ico
	// TODO: If this will be used in production and not statically hosted
	// then in memory caching of the image should be implemented here.
	http.Redirect(w, r, "/static/img/favicon.ico", 301)
}
