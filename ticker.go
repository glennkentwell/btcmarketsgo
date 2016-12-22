package btcmarketsgo

import "time"

//Ticker sets up a ticker to run the function provided every duration specified, quits when true is passed through quit chan
func (c BTCMarketsClient) Ticker(fn func(TickResponse, error), d time.Duration, quitChan chan bool) {
	go func() {
		for quit := false; !quit; {
			select {
			case quit = <-quitChan:
				break
			default:
				fn(c.Tick())
				time.Sleep(d)
			}
		}
	}()
}
