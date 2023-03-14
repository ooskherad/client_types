package client_type_worker

import (
	"encoding/json"
	"github.com/spf13/cast"
	"stock/models"
	"stock/stock_info_tse/client_types"
)

func SaveClientTypeInInputData(noInput ...interface{}) {
	clientTypesData := client_types.GetClientAll()
	var inputData []models.InputData
	for stockId, data := range clientTypesData {
		clientTypeModel := make(map[string]interface{})
		clientTypeModel["StockId"] = cast.ToInt64(stockId)
		clientTypeModel["NumberOfRealInBuy"] = data.NumberOfRealInBuy
		clientTypeModel["NumberOfLegalInBuy"] = data.NumberOfLegalInBuy
		clientTypeModel["NumberOfLegalInSell"] = data.NumberOfLegalInSell
		clientTypeModel["NumberOfRealInSell"] = data.NumberOfRealInSell
		clientTypeModel["VolumeOfRealInBuy"] = data.VolumeOfRealInBuy
		clientTypeModel["VolumeOfRealInSell"] = data.VolumeOfRealInSell
		clientTypeModel["VolumeOfLegalInSell"] = data.VolumeOfLegalInSell
		clientTypeModel["VolumeOfLegalInBuy"] = data.VolumeOfLegalInBuy

		inputDataData, _ := json.Marshal(clientTypeModel)
		inputData = append(inputData, models.InputData{
			Data:     inputDataData,
			Url:      client_types.ClientTypeAllUrl,
			Status:   1,
			DataType: 11,
		})
	}
	models.InputData{}.DB().Create(&inputData)
}
