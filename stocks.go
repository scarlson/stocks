package stocks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	url = "http://finance.google.com/finance/info?client=ig&q="
)

type quote struct {
	id        string    `json:"id"`
	Symbol    string    `json:"t"`
	Exchange  string    `json:"e"`
	Price     string    `json:"l"`
	price_fix string    `json:"l_fix"`
	price_cur string    `json:"l_cur"`
	s         int       `json:"s"`
	ltt       time.Time `json:"ltt"`
	lt        time.Time `json:"lt"`
	Change    string    `json:"c"`
	ChangePct string    `json:"cp"`
	cp_fix    float64   `json:"cp_fix"`
	ccol      string    `json:"ccol"`
}

func Quote(exchange, symbol string) (*quote, error) {
	u := url + exchange + ":" + symbol
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := string(body)[5:]
	str = str[:len(str)-2]
	q := &quote{}
	err = json.Unmarshal([]byte(str), q)
	if err != nil {
		return nil, err
	}
	return q, nil
}
