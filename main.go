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
	helper.Job(helper.JobTime{H: 9, M: 00, S: 00}, helper.JobTime{H: 12, M: 30, S: 00}, 3*time.Second, client_type_services.GetAndSaveInDb)
}
