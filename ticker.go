package btcmarketsgo

import "time"

//Ticker sets up a ticker to run the function provided every duration specified
func (c BTCMarketsClient) Ticker(fn func(TickResponse, error), d time.Duration) {
	go func() {
		for {
			fn(c.Tick())
			time.Sleep(d)
		}
	}()
}
