package main

import (
	"fmt"

	"github.com/RyanCarrier/btcmarketsgo"
	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
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
	got, err := client.GetOpenOrders()
	log.Info("Open orders output:")
	print(got, err)
}

/*
func main() {
	got, err := client.CreateBuyOrder(99900000000, 99900000000)
	print(got, err)
	print(client.OrdersDetails(got.ID))
	print(client.Tick())
	print(client.OrderBook())
	print(client.Trades())
	print(client.CancelOrder(got.ID))
	print(client.OrdersDetails(got.ID))

	//Ticker example
	quit := make(chan bool)
	client.Ticker(func(tr btcmarketsgo.TickResponse, err error) {
		fmt.Printf("%+v\n", tr)
	}, time.Second, quit)
	log.Info("quiting after 50 seconds")
	time.Sleep(time.Second * 5 * 10)
	quit <- true
	log.Info("quit")
}*/

func print(got interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	config := spew.NewDefaultConfig()
	config.Indent = "\t"
	config.Dump(got)
}
