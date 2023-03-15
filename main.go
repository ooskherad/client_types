package main

import (
	"fmt"
	"stock/helper"
	"stock/services/database"
	"stock/workers/client_type_worker"
	"time"
)

func main() {
	err := database.CreateDBConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	//con, err := database.GetDatabaseConnection()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = con.AutoMigrate(&models.Stock{}, &models.ClientType{}, &models.InputData{})
	//if err != nil {
	//	println(err.Error())
	//}
	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 30, S: 00}, 30*time.Second, client_type_worker.SaveClientTypes)
	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 30, S: 00}, 2*time.Second, client_type_worker.SaveClientTypeInInputData)
}
