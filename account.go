package btcmarketsgo

import (
	"encoding/json"
	"errors"
	"strings"

	ccg "github.com/RyanCarrier/cryptoclientgo"
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

//GetPrimaryCurrencies gets the account currencies available
func (c BTCMarketsClient) GetPrimaryCurrencies() ([]string, error) {
	return c.PrimaryCurrencies, nil
}

//GetSecondaryCurrencies gets the account currencies available
func (c BTCMarketsClient) GetSecondaryCurrencies() ([]string, error) {
	return c.SecondaryCurrencies, nil
}

//GetBalances gets the account balances
func (c BTCMarketsClient) GetBalances() (ccg.AccountBalances, error) {
	URI := "/account/balance"
	got, err := c.signAndGet(URI)
	if err != nil {
		log.Error("Error getting balance", err)
		return ccg.AccountBalances{}, err
	}
	var br BalancesResponse
	err = json.Unmarshal(got, &br)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	ab := ccg.AccountBalances{}

	for _, b := range br {
		ab = append(ab, ccg.AccountBalance{
			Currency:         b.Currency,
			AvailableBalance: b.Balance,
			TotalBalance:     b.Balance + b.PendingFunds,
		})
	}
	return ab, err
}

//GetBalance gets the balance of a single currency
func (c BTCMarketsClient) GetBalance(currency string) (ccg.AccountBalance, error) {
	got, err := c.GetBalances()
	if err != nil {
		return ccg.AccountBalance{}, err
	}
	currency = strings.ToUpper(strings.TrimSpace(currency))
	for _, b := range got {
		if currency == b.Currency {
			return b, nil
		}
	}
	return ccg.AccountBalance{}, errors.New("Currency " + currency + " not found")
}
