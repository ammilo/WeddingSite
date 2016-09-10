package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

// StaticController handles main static functions
func StaticController(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]
	fmt.Printf("Static request: %v\n", filename)
	if strings.Contains(r.URL.Path, "..") {
		http.NotFound(w, r)
	}

	http.ServeFile(w, r, r.URL.Path[1:])
}
