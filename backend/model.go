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

// TempCandle represents and item from the temp candle slice builiding the candles
type TempCandle struct {
	Symbol     string
	OpenTime   time.Time
	CloseTime  time.Time
	OpenPrice  float64
	ClosePrice float64
	HighPrice  float64
	LowPrice   float64
	volume     float64
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

// Data to write to clients  connected
type BroadcastMessaage struct {
	UpdateType UpdateType `json:"updateType"`
	Candle     *Candle    `json:"candle"`
}

type UpdateType string

const (
	Live   UpdateType = "live"  // Real time ongoing candle
	Closed UpdateType = "close" // Past candle. Already cloesd
)

// Converts a tempCodanlde into a candle
func (tc *TempCandle) toCaandle() *Candle {
	return &Candle{
		Symbol:    tc.Symbol,
		Open:      tc.OpenPrice,
		Close:     tc.ClosePrice,
		High:      tc.HighPrice,
		Low:       tc.LowPrice,
		Timestamp: tc.CloseTime,
	}
}
