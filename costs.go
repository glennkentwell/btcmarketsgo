package btcmarketsgo

import ccg "github.com/RyanCarrier/cryptoclientgo"

func (c BTCMarketsClient) GetTransactionCost(CurrencyFrom, CurrencyTo string) (ccg.Cost, error) {}

func (c BTCMarketsClient) GetWithdrawCost(Currency string) (ccg.Cost, error) {}

func (c BTCMarketsClient) GetDepositCost(Currency string) (ccg.Cost, error) {}
