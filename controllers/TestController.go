package controllers

import (
	"fmt"
	"net/http"

	"github.com/bytetrip/cordoctet/utils"
)

// TestController handles main homepage functions
func TestController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(utils.MakeUUID())
}
