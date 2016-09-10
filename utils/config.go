package utils

import (
	"encoding/json"
	"io/ioutil"
)

//Config type holds configs.
type Config struct {
	StartupMessage string
	ListenPort     string
}

// LoadConfig attempts to load config.json in the base directory
func (c *Config) LoadConfig() {
	b, e := ioutil.ReadFile("./config.json")
	CheckErr(e)
	json.Unmarshal(b, &c)
}
