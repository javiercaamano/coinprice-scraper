package models

import "time"

type RawMarketData struct {
	MarketData
	StartTime time.Time
	Value     float64
	Volume    float64
}

type Symbol struct {
	ProductID       string
	RawBase         string
	NormalizedBase  string
	RawQuote        string
	NormalizedQuote string
}

type OHLCMarketData struct {
	MarketData
	StartTime  time.Time
	EndTime    time.Time
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	Volume     float64
}

type MarketData struct {
	Source        string
	BaseCurrency  string
	QuoteCurrency string
}
