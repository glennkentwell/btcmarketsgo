package btcmarketsgo

import (
	"encoding/json"
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
)

//BalanceResponse is a single balance response
type BalanceResponse struct {
	Balance      int64  `json:"balance"`
	PendingFunds int64  `json:"pendingFunds"`
	Currency     string `json:"currency"`
}

//BalancesResponse is the response from requestiong balances
type BalancesResponse []BalanceResponse

//GetBalances gets the account balances
func (c BTCMarketsClient) GetBalances() (BalancesResponse, error) {
	URI := "/account/balance"
	got, err := c.signAndGet(URI)
	if err != nil {
		log.Error("Error getting balance", err)
		return BalancesResponse{}, err
	}
	var br BalancesResponse
	err = json.Unmarshal(got, &br)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return br, err
}

//GetBalance gets the balance of a single currency
func (c BTCMarketsClient) GetBalance(currency string) (BalanceResponse, error) {
	got, err := c.GetBalances()
	if err != nil {
		return BalanceResponse{}, err
	}
	currency = strings.ToUpper(strings.TrimSpace(currency))
	for _, b := range got {
		if currency == b.Currency {
			return b, nil
		}
	}
	return BalanceResponse{}, errors.New("Currency " + currency + " not found")
}
