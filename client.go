package btcmarketsgo

//BTCMarketsClient is primary struct for interacting with the API
type BTCMarketsClient struct {
	signiture string
}

//NewClient gets a new BTCMarketsClient
func NewClient(signiture string) *BTCMarketsClient {
	return &BTCMarketsClient{signiture: signiture}
}
