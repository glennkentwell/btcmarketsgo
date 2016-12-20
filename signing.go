package btcmarketsgo

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"strconv"
	"time"
)

func (c BTCMarketsClient) build(URI, body string) (int64, string) {
	now := time.Now().Unix() * 1000 //milliseconds
	return now, c.sign(URI + "\n" + strconv.FormatInt(now, 10) + "\n" + body)
}

func (c BTCMarketsClient) sign(message string) string {
	mac := hmac.New(sha512.New, c.decodedSecret)
	mac.Write([]byte(message))
	data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	//log.Info("Message to sign\n", message, "Signed message\n", data)
	return data
}
