package market_watch

import (
	"stock/crawler/services/http"
	"time"
)

const (
	MarketWatchUrl = "http://www.tsetmc.com/tsev2/data/MarketWatchPlus.aspx?h=0&r=0"
)

var (
	MarketWatchString string
	MarketWatchData   map[string]MarketWatchModel
)

type MarketWatchModel struct {
	Orders      Orders
	StockInfo   StockInfo
	StockPrices StockPrices
}
type StockPrices struct {
	Time               time.Time
	PriceFirst         int
	PriceMin           int
	PriceMax           int
	PriceYesterday     int
	PriceClose         int
	PriceClosePercent  float32
	PriceLastPercent   float32
	PriceLast          int
	TransactionCount   int
	TransactionVolume  int
	TransactionValue   int
	AuthorizedPriceMin int
	AuthorizedPriceMax int
}
type StockInfo struct {
	NameFa            string
	SymbolDigitCode12 string
	SymbolName        string
	TotalStockNumber  int64
	IndustryGroupCode string
	MonthAverage      int64
	EPS               int
	PE                float32
	BaseVolume        int
}

type Orders struct {
	R1                  OrderParts
	R2                  OrderParts
	R3                  OrderParts
	R4                  OrderParts
	R5                  OrderParts
	TotalBuyNumber      int
	TotalSellNumber     int
	TotalBuyVolume      int
	TotalSellVolume     int
	TotalBuyValue       int
	TotalSellValue      int
	MaximumVolumeOfRows int
	BuyAvgPrice         float32
	SellAvgPrice        float32
	BuyAvgPricePercent  float32
	SellAvgPricePercent float32
}
type OrderParts struct {
	Sell OrderItems
	Buy  OrderItems
}
type OrderItems struct {
	Number              int
	Price               int
	Volume              int
	PricePercent        float32
	IsInAuthorizedRange bool
}

func GetMarketWatch() map[string]MarketWatchModel {
	MarketWatchString, _ = http.GetAndString(MarketWatchUrl)
	modelMarketWatch(splitMarketWatch())
	return MarketWatchData
}
