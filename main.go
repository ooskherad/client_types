package main

import (
	"fmt"
	"stock/helper"
	"stock/models"
	"stock/services/client_type_services"
	"stock/services/database"
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
	err = con.AutoMigrate(&models.Stock{}, &models.ClientType{})
	if err != nil {
		println(err.Error())
	}

	client_type_services.GetAndSaveInDb()
	helper.Job(helper.JobTime{H: 00, M: 11, S: 00}, helper.JobTime{H: 00, M: 11, S: 10}, 3*time.Second, client_type_services.GetAndSaveInDb)
}
