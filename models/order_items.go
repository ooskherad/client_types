package models

import (
	"gorm.io/gorm"
	"log"
	"stock/services/database"
	"time"
)

type OrderItems struct {
	ID                   uint `gorm:"primarykey, AUTO_INCREMENT"`
	Stock                Stock
	StockId              uint `json:"stock_id"`
	RowNumber            int
	BNumber              int
	BPrice               int
	BVolume              int
	BPricePercent        float32
	BIsInAuthorizedRange bool
	SNumber              int
	SPrice               int
	SVolume              int
	SPricePercent        float32
	SIsInAuthorizedRange bool
	CreatedAt            time.Time
}

func (model OrderItems) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}
