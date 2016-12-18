package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/RyanCarrier/btcmarketsgo"
)

var public string
var private string

const apiPath = "api.secret"

func init() {
	getKeys()
}

func main() {
	client, err := btcmarketsgo.NewClient(public, private)
	if err == nil {
		got, err := client.Tick()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", got)
	}
}

func getKeys() {
	if _, err := os.Stat(apiPath); err == nil {
		getKeysFile()
	} else {
		getKeysStdin()
	}
}

func getKeysFile() {
	all, err := ioutil.ReadFile(apiPath)
	if err != nil {
		log.Error("error reading file", err)
		getKeysStdin()
		return
	}
	keys := strings.Fields(strings.TrimSpace(string(all)))
	if len(keys) != 2 {
		log.Error("fix your api.secret file with public, then private. Nothing else.")
		getKeysStdin()
	} else {
		public, private = keys[0], keys[1]
	}
}

func getKeysStdin() {
	fmt.Println("Public API key:")
	fmt.Scanln(&public)
	fmt.Println("Private API key:")
	fmt.Scanln(&private)
}
