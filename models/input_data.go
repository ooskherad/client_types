package models

import (
	"gorm.io/gorm"
	"log"
	"stock/infrastructure/helper/data_types"
	"stock/infrastructure/services/database"
	"time"
)

type InputData struct {
	Id        uint `gorm:"primarykey, AUTO_INCREMENT"`
	Data      data_types.JSON
	Url       string
	Status    int
	DataType  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model InputData) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}

func (model InputData) BulkUpdateVerifyStatus(models []InputData) {
	var ids []uint
	for _, inputData := range models {
		ids = append(ids, inputData.Id)
	}
	model.DB().Model(model).Where("id in ?", ids).UpdateColumn("status", 2)
}
