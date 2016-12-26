package btcmarketsgo

import (
	"errors"

	ccg "github.com/RyanCarrier/cryptoclientgo"
)

//GetTransactionCost gets the cost of a specified transaction
func (c BTCMarketsClient) GetTransactionCost(CurrencyFrom, CurrencyTo string) (ccg.Cost, error) {
	fi := lookupIndex(CurrencyFrom)
	ti := lookupIndex(CurrencyTo)
	if fi < 0 || ti < 0 {
		return ccg.Cost{}, errors.New("Could not find the currencies")
	}
	return ccg.Cost{Percent: TradeFees[fi][ti]}, nil
}

//GetWithdrawCost gets the cost of withdrawing from specified currency
func (c BTCMarketsClient) GetWithdrawCost(Currency string) (ccg.Cost, error) {
	if i := lookupIndex(Currency); i >= 0 {
		return ccg.Cost{Flat: WithdrawFees[i]}, nil
	}
	return ccg.Cost{}, errors.New("Could not find currency")
}

//GetDepositCost gets the cost of depositing to specified currency
func (c BTCMarketsClient) GetDepositCost(Currency string) (ccg.Cost, error) {
	if i := lookupIndex(Currency); i >= 0 {
		return ccg.Cost{Flat: DepositFees[i]}, nil
	}
	return ccg.Cost{}, errors.New("Could not find currency")
}
