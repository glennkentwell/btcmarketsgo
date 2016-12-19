package main

import (
	"fmt"
	"time"

	"github.com/RyanCarrier/btcmarketsgo"
	log "github.com/Sirupsen/logrus"
)

var public string
var private string
var client *btcmarketsgo.BTCMarketsClient

const apiPath = "api.secret"

func init() {
	public, private = getKeys()
}

func init() {
	var err error
	client, err = btcmarketsgo.NewClient(public, private)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	print(client.CreateBuyOrder(99900000000, 99900000000))

	//print(client.Tick())
	//print(client.OrderBook())
	//print(client.Trades())
	client.Ticker(func(tr btcmarketsgo.TickResponse, err error) {
		fmt.Printf("%+v\n", tr)
	}, time.Second*10)
	for {
	}
}

func print(got interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", got)
}
