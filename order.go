package btcmarketsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

//OrderStatuses is the current available order statuses
//const OrderStatuses = []string{"New", "Placed", "Failed", "Error", "Cancelled",	"Partially Canceled", "Fully Matched", "Partially Matched"}

//OrderRequest is an order request struct for parsing into json
type OrderRequest struct {
	Currency        string
	Instrument      string
	Price           int64
	Volume          int64
	OrderSide       string
	OrderType       string
	ClientRequestID string
}

//OrderResponse is the response from submitting an order
type OrderResponse struct {
	Success         bool
	ErrorCode       int
	ErrorMessage    string
	ID              int
	ClientRequestID string
}

//CreateOrder creates an order at specified price and volume
func (c BTCMarketsClient) createOrder(Price, Volume int64, Buy bool) (OrderResponse, error) {
	client := http.Client{}
	URL := c.Domain + "/order/create"
	or := OrderRequest{
		Currency:        c.Currency,
		Instrument:      c.Instrument,
		Price:           Price,
		Volume:          Volume,
		OrderSide:       "Bid",
		OrderType:       "Limit",
		ClientRequestID: "thisdoesntneedtobeset",
	}
	body, err := json.Marshal(or)
	if err != nil {
		return OrderResponse{}, err
	}
	fmt.Println("Posting\n", string(body))
	now, signature := c.build(URL, string(body))
	fmt.Println("String signed\n", string(signature))
	req, err := http.NewRequest("POST", URL, bytes.NewReader(body))
	if err != nil {
		return OrderResponse{}, errors.New("Error creating new Request;" + err.Error())
	}
	c.setupHeaders(req, now, string(signature))
	response, err := client.Do(req)
	if err != nil {
		return OrderResponse{}, errors.New("Error doing request;" + err.Error())
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return OrderResponse{}, errors.New("Error reading response;" + err.Error())
	}
	if response.StatusCode/100 != 2 {
		log.Error("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode) + "\n" + string(body))
		return OrderResponse{}, errors.New("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode))
	}
	var orderR OrderResponse
	err = json.Unmarshal(body, &orderR)
	if err != nil {
		return orderR, errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(body))
	}
	return orderR, nil
}

//CreateBuyOrder creates a buy order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateBuyOrder(Price, Volume int64) (OrderResponse, error) {
	return c.createOrder(Price, Volume, true)
}

/*
//CreateSellOrder creates a sell order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateSellOrder(Price, Volume int64)  (OrderResponse,error)  {
	return createOrder(Price, Volume, false)
}
*/
