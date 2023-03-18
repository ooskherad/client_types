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
	StockLastPrice      int
	CreatedAt           time.Time
}

func (model ClientType) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}

func (model ClientType) LastRecordOfClientType() []ClientType {
	query := "WITH ranked_messages AS (" +
		"    SELECT m.*, ROW_NUMBER() OVER (PARTITION BY stock_id ORDER BY created_at DESC) AS rn" +
		"    FROM client_types AS m" +
		") " +
		"SELECT * FROM ranked_messages WHERE rn = 1"
	var clientTypes []ClientType
	err := model.DB().Raw(query).Scan(&clientTypes)
	if err != nil {
		log.Println(err)
	}
	return clientTypes
}
