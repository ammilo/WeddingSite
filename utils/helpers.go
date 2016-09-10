package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// CheckErr returns nothing, but will panic if an error is passed to it.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// MakeUUID returns a UUID as a string.
func MakeUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	CheckErr(err)
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return strings.ToLower(uuid)
}
