package main

import (
	"fmt"

	"github.com/scarlson/stocks"
)

func main() {
	goog, err := stocks.Quote("nasdaq", "goog")
	if err != nil {
		panic(err)
	}
	fmt.Println(goog.Symbol, goog.Price, goog.Change)
}
