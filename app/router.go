package main

import (
	"net/http"

	"github.com/bytetrip/cordoctet/controllers"
)

func registerControllers() {

	// Specific handlers...
	http.HandleFunc("/test", controllers.TestController)
	http.HandleFunc("/favicon.ico", controllers.FaviconController)

	// All the rest...
	http.HandleFunc("/static/", controllers.StaticController)
	http.HandleFunc("/", controllers.RootController)

}
