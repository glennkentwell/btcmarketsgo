package btcmarketsgo

import (
	"encoding/json"
	"errors"

	ccg "github.com/RyanCarrier/cryptoclientgo"
)

//WithdrawRequest is the request built when attempting to withdraw
type WithdrawRequest struct {
	Amount   int64  `json:"amount"`
	Address  string `json:"address"`
	Currency string `json:"currency"`
}

//WithdrawResponse is the response recieved when requesting to withdraw
type WithdrawResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Status       string
}

//GetPrimaryCurrencyDepositAddress gets the deposit address for a digital currency
func (c BTCMarketsClient) GetPrimaryCurrencyDepositAddress(Currency string) (ccg.CurrencyAddress, error) {
	if i := lookupIndex(Currency); i >= 0 {
		return ccg.CurrencyAddress{DepositAddress: c.Addresses[i].Address}, nil
	}
	return ccg.CurrencyAddress{}, errors.New("Could not find currency")
}

//Withdraw withdraws the specified currency (and amount) to the specified BTC address.
// amount is *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) withdraw(amount int64, to string, currency string) (WithdrawResponse, error) {
	URI := "/fundtransfer/withdrawCrypto"
	wr := WithdrawRequest{amount, to, currency}
	got, err := c.signAndPost(URI, wr)
	var response WithdrawResponse
	err = json.Unmarshal(got, &response)
	if err != nil {
		err = errors.New("Error unmarshaling response;" + err.Error() + "\n" + string(got))
	}
	return response, err
}

//WithdrawCurrency withdraws the specified currency to the specified address
// amount is *10^-8, as specified in the BTCMarkets API;
// ie: $12.34 = 1,234,000,000; 12.34BTC=1,234,000,000
func (c BTCMarketsClient) WithdrawCurrency(Currency, To string, Amount int64) error {
	got, err := c.withdraw(Amount, To, Currency)
	if err != nil {
		return err
	}
	if got.Success {
		return nil
	}
	return errors.New("Withdraw fail; " + got.ErrorMessage)
}
