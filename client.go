package btcmarketsgo

import (
	"encoding/base64"

	log "github.com/Sirupsen/logrus"
)

//DefaultDomain is the default api domain
const DefaultDomain = "https://api.btcmarkets.net"

//DefaultCurrency is the default currency chosen for a new client
const DefaultCurrency = "AUD"

//DefaultInstrument is the default currency chosen for a new client
const DefaultInstrument = "BTC"

//BTCMarketsClient is primary struct for interacting with the API
type BTCMarketsClient struct {
	Public string
	//private       string
	decodedSecret []byte
	Domain        string
	Currency      string
	Instrument    string
}

//NewClient gets a new BTCMarketsClient
func NewClient(public, secret, domain, currency, instrument string) (*BTCMarketsClient, error) {
	data, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Fatal("error:", err)
		return nil, err
	}
	return &BTCMarketsClient{
		Public:        public,
		decodedSecret: data,
		Domain:        domain,
		Currency:      currency,
		Instrument:    instrument,
	}, nil
}

//NewDefaultClient gets a new client with default settings.
func NewDefaultClient(public, secret string) (*BTCMarketsClient, error) {
	return NewClient(public, secret, DefaultDomain, DefaultCurrency, DefaultInstrument)
}
