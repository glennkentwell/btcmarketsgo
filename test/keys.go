package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func getKeys() (string, string) {
	if _, err := os.Stat(apiPath); err == nil {
		return getKeysFile()
	}
	return getKeysStdin()

}

func getKeysFile() (string, string) {
	all, err := ioutil.ReadFile(apiPath)
	if err != nil {
		log.Error("error reading file", err)
		return getKeysStdin()
	}
	keys := strings.Fields(strings.TrimSpace(string(all)))
	if len(keys) != 2 {
		log.Error("fix your api.secret file with public, then private. Nothing else.")
		return getKeysStdin()
	}
	return keys[0], keys[1]

}

func getKeysStdin() (string, string) {
	var public, private string
	fmt.Println("Public API key:")
	fmt.Scanln(&public)
	fmt.Println("Private API key:")
	fmt.Scanln(&private)
	return public, private
}
