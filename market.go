package btcmarketsgo

import (
	"encoding/json"
	"strconv"
	"time"
)

//TickResponse is the response recieved when requesting market tick
type TickResponse struct {
	BestBid    float64
	BestAsk    float64
	LastPrice  float64
	Currency   string
	Instrument string
	Timestamp  int64
	Volume24h  float64
}

//Tick get current tick details
func (c BTCMarketsClient) Tick() (tr TickResponse, err error) {
	all, err := getBody(c.Domain + "/market/BTC/" + c.Currency + "/tick")
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}

//OrderBookResponse is the response recieved when requesting the order book
type OrderBookResponse struct {
	Currency   string
	Instrument string
	Timestamp  int64
	Asks       [][]float64
	Bids       [][]float64
}

//OrderBook gets the current orderbook
func (c BTCMarketsClient) OrderBook() (obr OrderBookResponse, err error) {
	all, err := getBody(c.Domain + "/market/BTC/" + c.Currency + "/orderbook")
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &obr)
	return
}

//TradeResponse is a single response of a trade
type TradeResponse struct {
	Tid    int64
	Amount float64
	Price  float64
	Date   int64
}

//TradesResponse is the trades returned from a trades request
type TradesResponse []TradeResponse

//Trades gets the current trades
func (c BTCMarketsClient) Trades() (TradesResponse, error) {
	return c.TradesSince(time.Time{})
}

//TradesSince gets the current trades since the specified time
func (c BTCMarketsClient) TradesSince(since time.Time) (tr TradesResponse, err error) {
	var all []byte
	if since.Equal(time.Time{}) {
		all, err = getBody(c.Domain + "/market/BTC/" + c.Currency + "/trades")
	} else {
		all, err = getBody(c.Domain + "/market/BTC/" + c.Currency + "/trades?since=" + strconv.FormatInt(since.Unix(), 10))
	}
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}
