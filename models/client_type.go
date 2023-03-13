package models

import (
	"gorm.io/gorm"
	"log"
	"stock/services/database"
	"time"
)

type ClientType struct {
	Id                  uint `gorm:"primarykey, AUTO_INCREMENT"`
	StockId             uint
	NumberOfRealInBuy   int
	NumberOfLegalInBuy  int
	NumberOfLegalInSell int
	NumberOfRealInSell  int
	VolumeOfRealInBuy   int
	VolumeOfRealInSell  int
	VolumeOfLegalInSell int
	VolumeOfLegalInBuy  int
	CreatedAt           time.Time
}

func (model ClientType) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}
