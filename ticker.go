package btcmarketsgo

import (
	"fmt"
	"time"
)

//Ticker sets up a ticker to run the function provided every duration specified, quits when true is passed through quit chan
func (c BTCMarketsClient) Ticker(fn func(TickResponse, error), d time.Duration, quitChan chan bool) {
	go func() {
		for {
			select {
			case quit := <-quitChan:
				if quit {
					fmt.Println("quit recieved", quit)
					return
				}
			default:
				fn(c.DefaultTick())
				time.Sleep(d)
			}
		}
	}()
}
