package market_watch_worker

import "time"

type InputData struct {
	Name    string `json:"Name,omitempty"`
	Message struct {
		Data   MarketWatchModel `json:"Data,omitempty"`
		Url    string           `json:"Url,omitempty"`
		Status int              `json:"Status,omitempty"`
	} `json:"Message,omitempty"`
	From      string    `json:"From,omitempty"`
	CreatedAt time.Time `json:"CreatedAt"`
}

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
	Identifier        string
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
