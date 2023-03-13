package client_type_services

import (
	"github.com/spf13/cast"
	"stock/client_types"
	"stock/models"
)

func GetAndSaveInDb(noInput ...interface{}) {
	clientTypesData := client_types.GetClientAll()
	var clientTypes []models.ClientType
	for stockId, data := range clientTypesData {
		clientTypeModel := models.ClientType{
			StockId:             cast.ToUint(stockId),
			NumberOfRealInBuy:   data.NumberOfRealInBuy,
			NumberOfLegalInBuy:  data.NumberOfLegalInBuy,
			NumberOfLegalInSell: data.NumberOfLegalInSell,
			NumberOfRealInSell:  data.NumberOfRealInSell,
			VolumeOfRealInBuy:   data.VolumeOfRealInBuy,
			VolumeOfRealInSell:  data.VolumeOfRealInSell,
			VolumeOfLegalInSell: data.VolumeOfLegalInSell,
			VolumeOfLegalInBuy:  data.VolumeOfLegalInBuy,
		}
		clientTypes = append(clientTypes, clientTypeModel)
	}
	models.ClientType{}.DB().Create(&clientTypes)
}
