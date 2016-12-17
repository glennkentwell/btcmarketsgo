package btcmarketsgo

//BTCMarketsClient is primary struct for interacting with the API
type BTCMarketsClient struct {
	signiture string
	api       string
}

//NewClient gets a new BTCMarketsClient
func NewClient(api, signiture string) *BTCMarketsClient {
	return &BTCMarketsClient{signiture: signiture, api: api}
}
