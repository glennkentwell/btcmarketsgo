package btcmarketsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

//OrderStatuses is the current available order statuses
//const OrderStatuses = []string{"New", "Placed", "Failed", "Error", "Cancelled",	"Partially Canceled", "Fully Matched", "Partially Matched"}

//OrderRequest is an order request struct for parsing into json
type OrderRequest struct {
	Currency        string `json:"currency"`
	Instrument      string `json:"instrument"`
	Price           int64  `json:"price"`
	Volume          int64  `json:"volume"`
	OrderSide       string `json:"orderSide"`       //Camel case
	OrderType       string `json:"ordertype"`       //the lowercase T is important...
	ClientRequestID string `json:"clientRequestId"` //Camel case
}

//OrderResponse is the response from submitting an order
type OrderResponse struct {
	Success         bool
	ErrorCode       int
	ErrorMessage    string
	ID              int
	ClientRequestID string
}

//CancelOrdersRequest is the struct used to request the cancelation of an order(s)
type CancelOrdersRequest struct {
	OrderIds []int `json:"orderIds"`
}

//CancelOrdersResponse is the response received when canceling an order(s)
type CancelOrdersResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Responses    []CancelOrderResponse
}

//CancelOrderResponse is the individual order cancelation response
type CancelOrderResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	ID           int
}

//CreateOrder creates an order at specified price and volume
func (c BTCMarketsClient) createOrder(Price, Volume int64, Buy bool) (OrderResponse, error) {
	URI := "/order/create"
	or := OrderRequest{
		Currency:        c.Currency,
		Instrument:      c.Instrument,
		Price:           Price,
		Volume:          Volume,
		OrderSide:       "Bid",
		OrderType:       "Limit",
		ClientRequestID: "1",
	}
	got, err := c.signAndPost(URI, or)
	var orderR OrderResponse
	err = json.Unmarshal(got, &orderR)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return orderR, err
}

//CancelOrder requests the cancelation of an order(s)
func (c BTCMarketsClient) CancelOrder(orderIDs ...int) (CancelOrdersResponse, error) {
	URI := "/order/cancel"
	cor := CancelOrdersRequest{OrderIds: orderIDs}
	got, err := c.signAndPost(URI, cor)
	var cancelOR CancelOrdersResponse
	err = json.Unmarshal(got, &cancelOR)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return cancelOR, err
}

//OrderHistoryRequest gets the users order history
type OrderHistoryRequest struct {
	Currency   string `json:"currency"`
	Instrument string `json:"instrument"`
	Limit      int    `json:"limit"`
	Since      int64  `json:"since,omitempty"`
}

//OrderHistoryResponse is the response returned when requesting the history of a user
type OrderHistoryResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Orders       []OrderHistorySingleResponse
}

//OrderHistorySingleResponse is a single order returned from a history request
type OrderHistorySingleResponse struct {
	ID              int64
	Currency        string
	Instrument      string
	OrderSide       string
	OrderType       string
	CreationTime    int64
	Status          string
	ErrorMessage    string
	Price           int64
	Volume          int64
	OpenVolume      int64
	ClientRequestID string
	Trades          []OrderHistoryTradeResponse
}

//OrderHistoryTradeResponse is a single trade from an order in a history request
type OrderHistoryTradeResponse struct {
	ID           int64
	CreationTime int64
	Description  string
	Price        int64
	Volume       int64
	Fee          int64
}

//OrderHistory gets the users order history
func (c BTCMarketsClient) OrderHistory() (OrderHistoryResponse, error) {
	return c.OrderHistorySince(0)
}

//OrderHistorySince gets the order history since specified time (Unix time in ms)
func (c BTCMarketsClient) OrderHistorySince(since int64) (OrderHistoryResponse, error) {
	return OrderHistoryResponse{}, nil
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

func (c BTCMarketsClient) signAndPost(URI string, i interface{}) ([]byte, error) {
	body, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	now, signature := c.sign(URI, string(body))
	URL := c.Domain + URI
	req, err := http.NewRequest("POST", URL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("Error creating new Request;" + err.Error())
	}
	c.setupHeaders(req, now, signature)
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error doing request;" + err.Error())
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading response;" + err.Error())
	}
	if response.StatusCode/100 != 2 {
		log.Error("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode) + "\n" + string(body))
		return nil, errors.New("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode))
	}
	return body, err
}
