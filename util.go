package btcmarketsgo

import (
	"io/ioutil"
	"net/http"
)

func getBody(request string) ([]byte, error) {
	resp, err := http.Get(request)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}
