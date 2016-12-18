package btcmarketsgo

import (
	"crypto/hmac"
	"crypto/sha512"
	"strconv"
	"time"
)

func (c BTCMarketsClient) build(URI, body string) (int64, []byte) {
	now := time.Now().Unix()
	return now, c.sign(URI + "\n" + strconv.FormatInt(now, 10) + "\n" + body)
}

func (c BTCMarketsClient) sign(message string) []byte {
	mac := hmac.New(sha512.New, c.decodedSecret)
	mac.Write([]byte(message))
	return mac.Sum(nil)
}
