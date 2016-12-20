package btcmarketsgo

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"strconv"
	"time"
)

func (c BTCMarketsClient) sign(URI, body string) (int64, string) {
	now := time.Now().Unix() * 1000 //milliseconds
	return now, c.hashEncode(URI + "\n" + strconv.FormatInt(now, 10) + "\n" + body)
}

func (c BTCMarketsClient) hashEncode(message string) string {
	mac := hmac.New(sha512.New, c.decodedSecret)
	mac.Write([]byte(message))
	data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return data
}
