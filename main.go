package main

import (
	"fmt"
	"stock/helper"
	"stock/models"
	"stock/services/database"
	"stock/workers/client_type_worker"
	"stock/workers/market_watch_worker"
	"time"
)

func main() {
	err := database.CreateDBConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	con, err := database.GetDatabaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = con.AutoMigrate(&models.Stock{}, &models.ClientType{}, &models.InputData{}, &models.StockPrices{}, &models.OrderItems{})
	if err != nil {
		println(err.Error())
	}
	market_watch_worker.SaveMarketWatchData()

	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 31, S: 00}, 3*time.Second, market_watch_worker.SaveMarketWatchData)
	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 31, S: 00}, 30*time.Second, client_type_worker.SaveClientTypes)
	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 30, S: 00}, 2*time.Second, client_type_worker.SaveClientTypeInInputData)
}
