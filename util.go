package btcmarketsgo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

func getBody(request string) ([]byte, error) {
	resp, err := http.Get(request)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (c BTCMarketsClient) sign(URI, body string) (int64, string) {
	now := time.Now().Unix() * 1000 //milliseconds
	return now, c.hashEncode(URI + "\n" + strconv.FormatInt(now, 10) + "\n" + body)
}

func (c BTCMarketsClient) hashEncode(message string) string {
	fmt.Println("SIGNING:\n" + message + "\nSIGNED")
	mac := hmac.New(sha512.New, c.decodedSecret)
	mac.Write([]byte(message))
	data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	fmt.Println(data + "\nSIGNED")
	return data
}

func (c BTCMarketsClient) setupHeaders(req *http.Request, timestamp int64, signature string) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.Public)
	req.Header.Set("timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("signature", signature)
}

func (c BTCMarketsClient) signAndPost(URI string, i interface{}) ([]byte, error) {
	body, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	now, signature := c.sign(URI, string(body))
	URL := c.Domain + URI
	req, err := http.NewRequest("POST", URL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("Error creating new Request;" + err.Error())
	}
	c.setupHeaders(req, now, signature)
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error doing request;" + err.Error())
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading response;" + err.Error())
	}
	if response.StatusCode/100 != 2 {
		log.Error("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode) + "\n" + string(body))
		return nil, errors.New("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode))
	}
	return body, err
}
