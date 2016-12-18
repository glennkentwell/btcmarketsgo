package btcmarketsgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//TickResponse is the response recieved when asking for market tick
type TickResponse struct {
	bestBid    float32
	bestAsk    float32
	lastPrice  float32
	currency   string
	instrument string
	timestamp  int64
	volume24h  float32
}

//Tick get current tick details
func (c BTCMarketsClient) Tick() (tr TickResponse, err error) {
	tr = TickResponse{}
	resp, err := http.Get(c.Domain + "/market/BTC/AUD/tick")
	if err != nil {
		return TickResponse{}, err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(all, &tr)
	return
}
