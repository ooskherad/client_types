package client_types

import (
	"stock/crawler/services/http"
)

const (
	ClientTypeAllUrl = "http://www.tsetmc.com/tsev2/data/ClientTypeAll.aspx"
)

var (
	ClientTypeAllString string
	ClientTypeAllData   map[string]ClientTypeAllModel
)

type ClientTypeAllModel struct {
	NumberOfRealInBuy   int
	NumberOfLegalInBuy  int
	NumberOfLegalInSell int
	NumberOfRealInSell  int
	VolumeOfRealInBuy   int
	VolumeOfRealInSell  int
	VolumeOfLegalInSell int
	VolumeOfLegalInBuy  int
}

func GetClientAll() map[string]ClientTypeAllModel {
	ClientTypeAllString, _ = http.GetAndString(ClientTypeAllUrl)
	modelClientAll(splitClientAll())
	return ClientTypeAllData
}
