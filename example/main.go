package main

import (
	"fmt"

	"github.com/RyanCarrier/btcmarketsgo"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

var client *btcmarketsgo.BTCMarketsClient

func init() {

	var err error
	client, err = btcmarketsgo.NewDefaultClient(btcmarketsgo.GetKeys("api.secret"))
	log.SetLevel(log.DebugLevel)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//got, err := client.GetOrderBook("BTC", "AUD")
	got, err := client.GetOrderBook("ETH", "BTC")
	log.Info("Open orders output:")
	print(got, err)
	//Ticker example
	/*quit := make(chan bool)
	client.Ticker(func(tr btcmarketsgo.TickResponse, err error) {
		fmt.Printf("%+v\n", tr)
	}, time.Second, quit)
	log.Info("quiting after 50 seconds")
	time.Sleep(time.Second * 5 * 10)
	quit <- true
	log.Info("quit")*/
}

func print(got interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	config := spew.NewDefaultConfig()
	config.Indent = "\t"
	config.Dump(got)
}
