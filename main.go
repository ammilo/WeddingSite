package main

import (
	"fmt"
	"net/http"

	"github.com/bytetrip/cordoctet/utils"
)

var config = new(utils.Config)

func main() {
	config.LoadConfig()
	fmt.Println(config.StartupMessage)

	registerControllers() //router.go

	http.ListenAndServe(config.ListenPort, nil)
}
