package util

import (
	"log"
	"testing"
)

func TestGetExternalIP(t *testing.T) {
	ip, err := GetExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(ip)
}
