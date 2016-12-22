package btcmarketsgo

import "time"

//Ticker sets up a ticker to run the function provided every duration specified, quits when true is passed through quit chan
func (c BTCMarketsClient) Ticker(fn func(TickResponse, error), d time.Duration, quit chan bool) {
	go func() {
		for {
			select {
			case val := <-quit:
				if val {
					return
				}
			default:
				fn(c.Tick())
				time.Sleep(d)
			}
		}
	}()
}
