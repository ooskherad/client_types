package client_type_worker

import (
	"encoding/json"
	"stock/models"
	"time"
)

func SaveClientTypes(noInput ...interface{}) {
	var clientTypesData []models.InputData
	models.InputData{}.DB().Where("status = ? and data_type = ?", 1, 11).Find(&clientTypesData)
	if len(clientTypesData) == 0 {
		return
	}

	lastClientTypeInDb := models.ClientType{}.LastRecordOfRecord()
	var clientTypeToSave []models.ClientType

	for _, clientTypeItem := range clientTypesData {
		find := false
		data := models.ClientType{}
		_ = json.Unmarshal(clientTypeItem.Data, &data)
		clientTypeItem.Status = 2
		clientTypeItem.VerifyAt = time.Now()

		for _, lastClientType := range lastClientTypeInDb {
			if data.StockId == lastClientType.StockId {
				if data.NumberOfRealInBuy != lastClientType.NumberOfRealInBuy ||
					data.NumberOfLegalInBuy != lastClientType.NumberOfLegalInBuy {
					clientTypeToSave = append(clientTypeToSave, data)
				}
				find = true
				break

			}
		}
		if !find {
			clientTypeToSave = append(clientTypeToSave, data)
		}
	}
	if len(clientTypeToSave) != 0 {
		models.ClientType{}.DB().Create(&clientTypeToSave)
	}
	models.InputData{}.BulkUpdateVerifyStatus(clientTypesData)
}
