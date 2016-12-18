package main

import (
	"fmt"

	"github.com/RyanCarrier/btcmarketsgo"
)

var public string
var private string

func init() {
	fmt.Println("Public API key:")
	fmt.Scanln(&public)
	fmt.Println("Private API key:")
	fmt.Scanln(&private)

}

func main() {
	client, err := btcmarketsgo.NewClient(public, private)
	if err == nil {
		fmt.Println(client.Tick())
	}
}
