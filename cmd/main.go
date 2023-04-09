package main

import (
	"fmt"
	"stock/pkg/data_processing/stock_info_tse/client_type_worker"
	"stock/pkg/data_processing/stock_info_tse/market_watch_worker"
	"stock/pkg/infrastructure/helper"
	"stock/pkg/infrastructure/services/database"
	"stock/pkg/models"

	"time"
)

func main() {
	database.Init()

	fmt.Println("connected to database")

	con, err := database.GetDatabaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = con.AutoMigrate(&models.Stock{}, &models.ClientType{}, &models.InputData{}, &models.StockPrices{}, &models.OrderItems{})
	if err != nil {
		println(err.Error())
	}
	done := make(chan bool)
	go helper.Job(helper.JobTime{H: 8, M: 59, S: 00}, helper.JobTime{H: 15, M: 00, S: 00}, 3*time.Second, market_watch_worker.SaveMarketWatchData)
	go helper.Job(helper.JobTime{H: 8, M: 59, S: 00}, helper.JobTime{H: 15, M: 00, S: 00}, 30*time.Second, client_type_worker.SaveClientTypes)
	fmt.Println("start running workers \t ctrl+c to break")
	<-done

}
