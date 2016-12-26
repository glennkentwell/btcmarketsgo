package btcmarketsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	ccg "github.com/RyanCarrier/cryptoclientgo"
)

const btcMin = int64(100000)
const multiplier = int64(100000000)

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
func (c BTCMarketsClient) createOrder(Price, Volume int64, Buy bool, Market bool) (OrderResponse, error) {
	if Volume < btcMin {
		return OrderResponse{}, errors.New(
			fmt.Sprint("Volume must be ", btcMin, " minimum (", strconv.FormatFloat(
				float64(btcMin)/float64(multiplier), 'f', 3, 64)+"BTC)",
			),
		)
	}
	URI := "/order/create"
	or := OrderRequest{
		Currency:        c.Currency,
		Instrument:      c.Instrument,
		Price:           Price,
		Volume:          Volume,
		ClientRequestID: "1",
	}
	if Buy {
		or.OrderSide = "Bid"
	} else {
		or.OrderSide = "Ask"
	}
	if Market {
		or.OrderType = "Market"
	} else {
		or.OrderType = "Limit"
	}
	got, err := c.signAndPost(URI, or)
	var orderR OrderResponse
	err = json.Unmarshal(got, &orderR)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return orderR, err
}

//CancelOrders requests the cancelation of an order(s)
func (c BTCMarketsClient) CancelOrders(orderIDs ...int) (CancelOrdersResponse, error) {
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

//CancelOrder cancels a single order
func (c BTCMarketsClient) CancelOrder(orderID int) error {
	got, err := c.CancelOrders(orderID)
	if err != nil {
		return err
	}
}

//OrderHistoryRequest gets the users order history
type OrderHistoryRequest struct {
	Currency   string `json:"currency"`
	Instrument string `json:"instrument"`
	Limit      int    `json:"limit"`
	Since      int64  `json:"since"`
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
func (c BTCMarketsClient) OrderHistory(limit int) (OrderHistoryResponse, error) {
	return c.OrderHistorySince(limit, 0)
}

//OrderHistorySince gets the order history since specified time (Unix time in ms)
func (c BTCMarketsClient) OrderHistorySince(limit int, since int64) (OrderHistoryResponse, error) {
	return c.orderHistory(limit, since, 1)
}

//mode;
//0 Open order history
//1 All order history
//2 Trade history
func (c BTCMarketsClient) orderHistory(limit int, since int64, mode int) (OrderHistoryResponse, error) {
	var URI string
	switch mode {
	case 0:
		URI = "/order/open"
		break
	case 1:
		URI = "/order/history"
		break
	case 2:
		URI = "/order/trade/history"
		break
	default:
		return OrderHistoryResponse{}, errors.New("mode somehow set incorrectly in private function")
	}
	ohr := OrderHistoryRequest{
		Currency:   c.Currency,
		Instrument: c.Instrument,
		Limit:      limit,
		Since:      since,
	}
	got, err := c.signAndPost(URI, ohr)
	var ohs OrderHistoryResponse
	err = json.Unmarshal(got, &ohs)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return ohs, err
}

//GetOpenOrders gets the current open orders
func (c BTCMarketsClient) GetOpenOrders() (ccg.OrdersDetails, error) {
	got, err := c.orderHistory(9999, 0, 1)
	return ccg.OrderDetails{}, nil
}

//TradeHistory gets the current trade history
func (c BTCMarketsClient) TradeHistory(limit int) (OrderHistoryResponse, error) {
	return c.TradeHistorySince(limit, 0)
}

//TradeHistorySince gets the current trade history since the time specified (Unix ms)
func (c BTCMarketsClient) TradeHistorySince(limit int, since int64) (OrderHistoryResponse, error) {
	return c.orderHistory(limit, since, 2)
}

//OrderDetailsRequest is the struct used to request the details for order(s)
type OrderDetailsRequest struct {
	OrderIds []int `json:"orderIds"`
}

//OrdersDetailsResponse is the response recieved from order details requests
type OrdersDetailsResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Orders       []OrderDetailsResponse
}

//OrderDetailsResponse is the details returned from a single order
type OrderDetailsResponse struct {
	ID           int
	Currency     string
	Instrument   string
	OrderSide    string
	OrderType    string
	CreationTime int64
	Status       string
	ErrorMessage string
	Price        int64
	Volume       int64
	OpenVolume   int64
	Trades       []OrderHistoryTradeResponse
}

//OrdersDetails gets the details of the specified orders
func (c BTCMarketsClient) OrdersDetails(orderIDs ...int) (OrdersDetailsResponse, error) {
	URI := "/order/detail"
	cor := OrderDetailsRequest{OrderIds: orderIDs}
	got, err := c.signAndPost(URI, cor)
	var odr OrdersDetailsResponse
	err = json.Unmarshal(got, &odr)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return OrdersDetailsResponse{}, err
}

//OrderDetails gets a single orders details
func (c BTCMarketsClient) OrderDetails(orderID int) (ccg.OrderDetails, error) {
	got, err := c.OrdersDetails(orderID)
	return ccg.OrderDetails{}, err
}

//CreateBuyOrder creates a buy order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateBuyOrder(Price, Volume int64) (OrderResponse, error) {
	return c.createOrder(Price, Volume, true, false)
}

//PlaceMarketOrder places a market order
func (c BTCMarketsClient) PlaceMarketOrder(CurrencyFrom, CurrencyTo string, amount int64) (ccg.OrderDetails, error) {

}

//PlaceLimitOrder places a limit order for the specified price, that is, the price and amount will be the trades.
func (c BTCMarketsClient) PlaceLimitOrder(CurrencyFrom, CurrencyTo string, amount, price int64) (ccg.OrderDetails, error) {

}

//GetOrderBook gets the orders for the relevant currencies
func (c BTCMarketsClient) GetOrderBook(CurrencyFrom, CurrencyTo string) (ccg.OrderBook, error) {

}

//GetRecentTrades gets most recent trades limited by historyAmount
func (c BTCMarketsClient) GetRecentTrades(CurrencyFrom, CurrencyTo string, historyAmount int) (ccg.RecentTrades, error) {

}

//CreateMarketBuyOrder creates a buy order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateMarketBuyOrder(Volume int64) (OrderResponse, error) {
	return c.createOrder(0, Volume, true, true)
}

//CreateSellOrder creates a sell order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateSellOrder(Price, Volume int64) (OrderResponse, error) {
	return c.createOrder(Price, Volume, false, false)
}

//CreateMarketSellOrder creates a sell order for the specified price and volume.
// Price and volume are both *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) CreateMarketSellOrder(Volume int64) (OrderResponse, error) {
	return c.createOrder(0, Volume, false, true)
}
