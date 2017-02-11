package btcmarketsgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//GetKeys gets the api keys from file specified or stdin if it doesn't exist.
func GetKeys(fileLocation string) (string, string, error) {
	if _, err := os.Stat(fileLocation); err == nil {
		return GetKeysFromFile(fileLocation)
	}
	return GetKeysFromStdin()
}

//GetKeysFromFile gets the keys from specified file
// format should be API key on first line, private key on second line.
func GetKeysFromFile(fileLocation string) (string, string, error) {
	all, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return "", "", errors.New("error reading file:" + err.Error())
	}
	keys := strings.Fields(strings.TrimSpace(string(all)))
	if len(keys) != 2 {
		return "", "", errors.New("fix your " + fileLocation + " file with public (API) key, then private key. Nothing else.")
	}
	return keys[0], keys[1], nil

}

//GetKeysFromStdin gets the API and private key from stdin
func GetKeysFromStdin() (string, string, error) {
	var public, private string
	fmt.Println("Public API key:")
	fmt.Scanln(&public)
	fmt.Println("Private key:")
	fmt.Scanln(&private)
	return public, private, nil
}
