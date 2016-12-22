package main

import (
	"fmt"
	"time"

	"github.com/RyanCarrier/btcmarketsgo"
	log "github.com/Sirupsen/logrus"
)

var client *btcmarketsgo.BTCMarketsClient

func init() {
	public, private := btcmarketsgo.GetKeys("api.secret")
	var err error
	//Could pass keys directly through, but abstracted for clarity
	client, err = btcmarketsgo.NewDefaultClient(public, private)
	//overwrite private here TODO (this won't need to be done if keys passed directly)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	print(client.CreateBuyOrder(99900000000, 99900000000))

	//print(client.Tick())
	//print(client.OrderBook())
	//print(client.Trades())

	if false {
		quit := make(chan bool)
		client.Ticker(func(tr btcmarketsgo.TickResponse, err error) {
			fmt.Printf("%+v\n", tr)
		}, time.Second*10, quit)
		time.Sleep(time.Second * 5 * 10)
		quit <- true
	}
	log.Info("quit")
}

func print(got interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", got)
}
