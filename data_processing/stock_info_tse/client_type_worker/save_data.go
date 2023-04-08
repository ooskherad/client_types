package client_type_worker

import (
	"encoding/json"
	"stock/models"
)

func SaveClientTypes(noInput ...interface{}) {
	var clientTypesData []models.InputData
	models.InputData{}.DB().Where("status = ? and data_type = ?", 1, 11).Order("created_at").Find(&clientTypesData)
	if len(clientTypesData) == 0 {
		return
	}

	lastClientTypeInDb := models.ClientType{}.LastRecordOfClientType()
	var clientTypeToSave []models.ClientType

	for _, clientTypeItem := range clientTypesData {
		find := false
		data := models.ClientType{}
		_ = json.Unmarshal(clientTypeItem.Data, &data)
		data.CreatedAt = clientTypeItem.CreatedAt
		clientTypeItem.Status = 2

		for i, lastClientType := range lastClientTypeInDb {
			if data.StockId == lastClientType.StockId {
				if data.NumberOfRealInBuy > lastClientType.NumberOfRealInBuy ||
					data.NumberOfLegalInBuy > lastClientType.NumberOfLegalInBuy ||
					data.NumberOfLegalInSell > lastClientType.NumberOfLegalInSell ||
					data.NumberOfRealInSell > lastClientType.NumberOfRealInSell {
					clientTypeToSave = append(clientTypeToSave, data)
					lastClientTypeInDb[i] = data
				}
				find = true
				break

			}
		}
		if !find {
			clientTypeToSave = append(clientTypeToSave, data)
			lastClientTypeInDb = append(lastClientTypeInDb, data)
		}
	}
	if len(clientTypeToSave) != 0 {
		models.ClientType{}.DB().Create(&clientTypeToSave)
	}
	models.InputData{}.BulkUpdateVerifyStatus(clientTypesData)
}
