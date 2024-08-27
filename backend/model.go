package main

import "time"

//Candle Struct represents a single OHLC(Open, High, Low, Close) candle

type Candle struct {
	Symbol    string    `json:"symbol"`
	Open      float64   `json:"open"`
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Timestamp time.Time `json:"timestamp"`
}

// structure of the data that comes from the Finnhub ws api
type FinnhubMessage struct {
	Data []TradeData `json:"data"`
	Type string      `json:"type"`
}

type TradeData struct {
	Close     []string `json:"c"`
	Price     float64  `json:"p"`
	Symbol    string   `json:"s"`
	Timestamp int64    `json:"t"`
	Volume    int      `json:"v"`
}
