package btcmarketsgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//TickResponse is the response recieved when asking for market tick
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
	tr = TickResponse{}
	resp, err := http.Get(c.Domain + "/market/BTC/AUD/tick")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}

//Tick get current tick details
func (c BTCMarketsClient) OrderBook() (tr TickResponse, err error) {
	tr = TickResponse{}
	resp, err := http.Get(c.Domain + "/market/BTC/AUD/tick")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}

//Tick get current tick details
func (c BTCMarketsClient) Trades() (tr TickResponse, err error) {
	tr = TickResponse{}
	resp, err := http.Get(c.Domain + "/market/BTC/AUD/tick")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}

//Tick get current tick details
func (c BTCMarketsClient) TradesSince(since time.Time) (tr TickResponse, err error) {
	tr = TickResponse{}
	resp, err := http.Get(c.Domain + "/market/BTC/AUD/tick")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}
