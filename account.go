package btcmarketsgo

import (
	"encoding/json"
	"errors"

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
