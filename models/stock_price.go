package models

import (
	"gorm.io/gorm"
	"log"
	"stock/infrastructure/services/database"
	"time"
)

type StockPrices struct {
	ID                uint `gorm:"primarykey, AUTO_INCREMENT"`
	Stock             Stock
	StockId           uint      `json:"stock_id"`
	PriceMin          int       `json:"price_min"`
	PriceMax          int       `json:"price_max"`
	PriceYesterday    int       `json:"price_yesterday"`
	PriceFirst        int       `json:"price_first"`
	PriceClose        int       `json:"price_close"`
	PriceLast         int       `json:"price_last"`
	TransactionCount  int       `json:"transaction_count"`
	TransactionVolume int       `json:"transaction_volume"`
	TransactionValue  int       `json:"transaction_value"`
	TransactionAt     time.Time `json:"transaction_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (model StockPrices) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}

func (model StockPrices) LastPriceOfStock() []StockPrices {
	query := "with prices as (select *, row_number() over (PARTITION BY stock_id order by transaction_at desc ) as row_number" +
		"                from stock_prices" +
		"                where transaction_at::date = current_date)" +
		" select *" +
		" from prices" +
		" where row_number = 1"
	var stockPrices []StockPrices
	err := model.DB().Raw(query).Scan(&stockPrices)
	if err != nil {
		log.Println(err)
	}
	return stockPrices
}
