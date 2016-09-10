package main

import "github.com/bytetrip/cordoctet/utils"

var config = new(utils.Config)

func init() {
	config.LoadConfig()

	registerControllers() //router.go

	//http.ListenAndServe(config.ListenPort, nil)

}
