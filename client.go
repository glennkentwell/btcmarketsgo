package btcmarketsgo

import (
	"encoding/base64"
	"log"
)

//Domain is the api base domain
const Domain = "https://api.btcmarkets.net"

//BTCMarketsClient is primary struct for interacting with the API
type BTCMarketsClient struct {
	Public        string
	private       string
	decodedSecret []byte
	Domain        string
}

//NewClient gets a new BTCMarketsClient
func NewClient(public, secret string) (*BTCMarketsClient, error) {
	data, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Fatal("error:", err)
		return nil, err
	}
	return &BTCMarketsClient{Public: public, private: secret, decodedSecret: data, Domain: Domain}, nil
}
