package client_types

import (
	"sort"
	"stock/helper"
	"strings"
)

func splitClientAll() (results [][]string) {
	if ClientTypeAllString == "" {
		return
	}
	for _, r := range strings.Split(ClientTypeAllString, ";") {
		items := strings.Split(r, ",")
		if len(items) == 9 {
			results = append(results, items)
		}
	}
	return
}

func modelClientAll(splitData [][]string) {
	if len(splitData) == 0 {
		return
	}

	sort.SliceStable(splitData, func(m, n int) bool {
		return splitData[m][0] < splitData[n][0]
	})
	clientTypeAllModel := make(map[string]ClientTypeAllModel, len(splitData))
	for i := 0; i < len(splitData); i++ {
		model := ClientTypeAllModel{
			NumberOfRealInBuy: helper.StringToInt(splitData[i][1]), NumberOfLegalInBuy: helper.StringToInt(splitData[i][2]),
			VolumeOfRealInBuy: helper.StringToInt(splitData[i][3]), VolumeOfLegalInBuy: helper.StringToInt(splitData[i][4]),
			NumberOfRealInSell: helper.StringToInt(splitData[i][5]), NumberOfLegalInSell: helper.StringToInt(splitData[i][6]),
			VolumeOfRealInSell: helper.StringToInt(splitData[i][7]), VolumeOfLegalInSell: helper.StringToInt(splitData[i][8])}

		clientTypeAllModel[splitData[i][0]] = model
	}
	ClientTypeAllData = clientTypeAllModel
}

func FindStockInClientAll(identifier string) ClientTypeAllModel {
	value, _ := ClientTypeAllData[identifier]
	return value
}
