package btcmarketsgo

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func getBody(request string) ([]byte, error) {
	resp, err := http.Get(request)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (c BTCMarketsClient) setupHeaders(req *http.Request, timestamp int64, signature string) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.Public)
	req.Header.Set("timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("signature", signature)
}
