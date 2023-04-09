package models

import (
	"gorm.io/gorm"
	"log"
	"stock/infrastructure/services/database"
	"time"
)

type Stock struct {
	ID                   uint   `gorm:"primarykey, AUTO_INCREMENT"`
	NameFa               string `json:"name_fa" gorm:"type:varchar(40); unique_index"`
	NameEn               string `json:"name_en" gorm:"type:varchar(40)"`
	CompanyDigitCode12   string `json:"company_digit_code12"  gorm:"type:varchar(12)"`
	SymbolDigitCode12    string `json:"symbol_digit_code12"  gorm:"type:varchar(12)"`
	SymbolDigit5         string `json:"symbol_digit5"  gorm:"type:varchar(5)"`
	CompanyDigit4        string `json:"company_digit4" gorm:"type:varchar(4)"`
	SymbolName           string `json:"symbol_name"`
	Market               string `json:"market"  gorm:"type:varchar(35)"`
	TableCode            int    `json:"table_code"`
	IndustryGroupCode    int    `json:"industry_group_code"`
	IndustrySubgroupCode int    `json:"industry_subgroup_code"`
	IndustryGroupName    string `json:"industry_group_name"  gorm:"type:varchar(40)"`
	IndustrySubgroupName string `json:"industry_subgroup_name"  gorm:"type:varchar(40)"`
	TotalStockNumber     int64
	MonthAverage         int64
	EPS                  int
	PE                   float32
	BaseVolume           int
	CreatedAT            time.Time
}

func (model Stock) DB() *gorm.DB {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
	}
	return db
}
