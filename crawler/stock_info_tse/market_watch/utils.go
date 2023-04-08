package market_watch

import (
	"github.com/spf13/cast"
	"sort"
	"stock/infrastructure/helper"
	"strings"
	"time"
	"unicode/utf8"
)

var y, m, d = time.Now().Date()

func splitMarketWatch() (info [][]string, orders [][]string) {
	if MarketWatchString == "" {
		return
	}
	MarketWatchString = strings.Replace(MarketWatchString, "@", ";", -1)
	for _, r := range strings.Split(MarketWatchString, ";") {
		items := strings.Split(r, ",")
		if len(items) == 25 || len(items) == 23 {
			info = append(info, items)
		} else if len(items) == 8 {
			orders = append(orders, items)
		}
	}
	return
}

func modelMarketWatch(splitInfo [][]string, splitOrders [][]string) {
	if len(splitInfo) == 0 {
		return
	}
	// StockInfo model
	marketWatchData := make(map[string]MarketWatchModel, len(splitInfo))
	for i := 0; i < len(splitInfo); i++ {
		var model MarketWatchModel
		info := StockInfo{
			// http://www.tsetmc.com/loader.aspx?ParTree=151311&i=65883838195688438
			//, MonthAverage: QTotTran5JAvg
			//KAjCapValCpsIdx= سهام شناور
			SymbolDigitCode12: splitInfo[i][1], SymbolName: splitInfo[i][2], NameFa: splitInfo[i][3],
			EPS: helper.StringToInt(splitInfo[i][14]), BaseVolume: helper.StringToInt(splitInfo[i][15]),
			IndustryGroupCode: splitInfo[i][18], TotalStockNumber: cast.ToInt64(splitInfo[i][21])}

		price := StockPrices{
			Time: convertStringToTile(splitInfo[i][4]), PriceFirst: helper.StringToInt(splitInfo[i][5]),
			PriceLast: helper.StringToInt(splitInfo[i][7]), PriceClose: helper.StringToInt(splitInfo[i][6]),
			TransactionCount: helper.StringToInt(splitInfo[i][8]), TransactionVolume: helper.StringToInt(splitInfo[i][9]), TransactionValue: helper.StringToInt(splitInfo[i][10]),
			PriceMin: helper.StringToInt(splitInfo[i][11]), PriceMax: helper.StringToInt(splitInfo[i][12]),
			PriceYesterday: helper.StringToInt(splitInfo[i][13]), AuthorizedPriceMax: helper.FloatStringToInt(splitInfo[i][19]),
			AuthorizedPriceMin: helper.FloatStringToInt(splitInfo[i][20])}

		// calculate Percents
		price.PriceClosePercent = float32(price.PriceClose-price.PriceYesterday) / float32(price.PriceYesterday) * 100
		price.PriceLastPercent = float32(price.PriceLast-price.PriceYesterday) / float32(price.PriceYesterday) * 100
		info.PE = helper.Divide(price.PriceClose, info.EPS)

		model.StockInfo = info
		model.StockPrices = price
		marketWatchData[splitInfo[i][0]] = model
	}

	// Orders table Model
	sort.SliceStable(splitOrders, func(m, n int) bool {
		return splitOrders[m][0] < splitOrders[n][0]
	})
	for i := 0; i < len(splitOrders); i++ {
		stockMarketWatchData, found := marketWatchData[splitOrders[i][0]]
		if !found {
			continue
		}
		orderRaw := OrderParts{
			Buy: OrderItems{
				Number: helper.StringToInt(splitOrders[i][3]), Price: helper.StringToInt(splitOrders[i][4]),
				Volume: helper.StringToInt(splitOrders[i][6])},
			Sell: OrderItems{
				Number: helper.StringToInt(splitOrders[i][2]), Price: helper.StringToInt(splitOrders[i][5]),
				Volume: helper.StringToInt(splitOrders[i][7])},
		}
		if orderRaw.Buy.Volume > stockMarketWatchData.Orders.MaximumVolumeOfRows {
			stockMarketWatchData.Orders.MaximumVolumeOfRows = orderRaw.Buy.Volume
		}
		if orderRaw.Sell.Volume > stockMarketWatchData.Orders.MaximumVolumeOfRows {
			stockMarketWatchData.Orders.MaximumVolumeOfRows = orderRaw.Sell.Volume
		}

		// calculate Percents
		if orderRaw.Sell.Price != 0 {
			orderRaw.Sell.PricePercent = float32(orderRaw.Sell.Price-stockMarketWatchData.StockPrices.PriceYesterday) / float32(stockMarketWatchData.StockPrices.PriceYesterday) * 100
			isPriceLowerThanMaxAuthPrice := orderRaw.Sell.Price <= stockMarketWatchData.StockPrices.AuthorizedPriceMax
			isPriceGreaterThanMinAuthPrice := orderRaw.Sell.Price >= stockMarketWatchData.StockPrices.AuthorizedPriceMin
			orderRaw.Sell.IsInAuthorizedRange = isPriceLowerThanMaxAuthPrice && isPriceGreaterThanMinAuthPrice
		}
		if orderRaw.Buy.Price != 0 {
			orderRaw.Buy.PricePercent = float32(orderRaw.Buy.Price-stockMarketWatchData.StockPrices.PriceYesterday) / float32(stockMarketWatchData.StockPrices.PriceYesterday) * 100
			isPriceLowerThanMaxAuthPrice := orderRaw.Buy.Price <= stockMarketWatchData.StockPrices.AuthorizedPriceMax
			isPriceGreaterThanMinAuthPrice := orderRaw.Buy.Price >= stockMarketWatchData.StockPrices.AuthorizedPriceMin
			orderRaw.Buy.IsInAuthorizedRange = isPriceLowerThanMaxAuthPrice && isPriceGreaterThanMinAuthPrice
		}

		if splitOrders[i][1] == "1" {
			stockMarketWatchData.Orders.R1 = orderRaw
		} else if splitOrders[i][1] == "2" {
			stockMarketWatchData.Orders.R2 = orderRaw
		} else if splitOrders[i][1] == "3" {
			stockMarketWatchData.Orders.R3 = orderRaw
		} else if splitOrders[i][1] == "4" {
			stockMarketWatchData.Orders.R4 = orderRaw
		} else if splitOrders[i][1] == "5" {
			stockMarketWatchData.Orders.R5 = orderRaw
		}
		marketWatchData[splitOrders[i][0]] = stockMarketWatchData

	}

	// calculate Orders data
	for identifier, stockMarketWatchData := range marketWatchData {
		orders := stockMarketWatchData.Orders

		stockMarketWatchData.Orders.TotalBuyNumber = orders.R1.Buy.Number + orders.R2.Buy.Number + orders.R3.Buy.Number + orders.R4.Buy.Number + orders.R5.Buy.Number
		stockMarketWatchData.Orders.TotalSellNumber = orders.R1.Sell.Number + orders.R2.Sell.Number + orders.R3.Sell.Number + orders.R4.Sell.Number + orders.R5.Sell.Number

		stockMarketWatchData.Orders.TotalBuyVolume = orders.R1.Buy.Volume + orders.R2.Buy.Volume + orders.R3.Buy.Volume + orders.R4.Buy.Volume + orders.R5.Buy.Volume
		stockMarketWatchData.Orders.TotalSellVolume = orders.R1.Sell.Volume + orders.R2.Sell.Volume + orders.R3.Sell.Volume + orders.R4.Sell.Volume + orders.R5.Sell.Volume

		stockMarketWatchData.Orders.TotalBuyValue = orders.R1.Buy.Volume*orders.R1.Buy.Price + orders.R2.Buy.Volume*orders.R2.Buy.Price + orders.R3.Buy.Volume*orders.R3.Buy.Price + orders.R4.Buy.Volume*orders.R4.Buy.Price + orders.R5.Buy.Volume*orders.R5.Buy.Price
		stockMarketWatchData.Orders.TotalSellValue = orders.R1.Sell.Volume*orders.R1.Sell.Price + orders.R2.Sell.Volume*orders.R2.Sell.Price + orders.R3.Sell.Volume*orders.R3.Sell.Price + orders.R4.Sell.Volume*orders.R4.Sell.Price + orders.R5.Sell.Volume*orders.R5.Sell.Price

		totalBuyWeight := helper.Divide(orders.R1.Buy.Volume, orders.R1.Buy.Number) + helper.Divide(orders.R2.Buy.Volume, orders.R2.Buy.Number) +
			helper.Divide(orders.R3.Buy.Volume, orders.R3.Buy.Number) + helper.Divide(orders.R4.Buy.Volume, orders.R4.Buy.Number) +
			helper.Divide(orders.R5.Buy.Volume, orders.R5.Buy.Number)

		stockMarketWatchData.Orders.BuyAvgPrice = helper.Divide(
			(helper.Divide(orders.R1.Buy.Volume, orders.R1.Buy.Number)*float32(orders.R1.Buy.Price))+
				(helper.Divide(orders.R2.Buy.Volume, orders.R2.Buy.Number)*float32(orders.R2.Buy.Price))+
				(helper.Divide(orders.R3.Buy.Volume, orders.R3.Buy.Number)*float32(orders.R3.Buy.Price))+
				(helper.Divide(orders.R4.Buy.Volume, orders.R4.Buy.Number)*float32(orders.R4.Buy.Price))+
				(helper.Divide(orders.R5.Buy.Volume, orders.R5.Buy.Number)*float32(orders.R5.Buy.Price)), totalBuyWeight)

		totalSellWeight := helper.Divide(orders.R1.Sell.Volume, orders.R1.Sell.Number) + helper.Divide(orders.R2.Sell.Volume, orders.R2.Sell.Number) +
			helper.Divide(orders.R3.Sell.Volume, orders.R3.Sell.Number) + helper.Divide(orders.R4.Sell.Volume, orders.R4.Sell.Number) +
			helper.Divide(orders.R5.Sell.Volume, orders.R5.Sell.Number)

		stockMarketWatchData.Orders.SellAvgPrice = helper.Divide(
			(helper.Divide(orders.R1.Sell.Volume, orders.R1.Sell.Number)*float32(orders.R1.Sell.Price))+
				(helper.Divide(orders.R2.Sell.Volume, orders.R2.Sell.Number)*float32(orders.R2.Sell.Price))+
				(helper.Divide(orders.R3.Sell.Volume, orders.R3.Sell.Number)*float32(orders.R3.Sell.Price))+
				(helper.Divide(orders.R4.Sell.Volume, orders.R4.Sell.Number)*float32(orders.R4.Sell.Price))+
				(helper.Divide(orders.R5.Sell.Volume, orders.R5.Sell.Number)*float32(orders.R5.Sell.Price)), totalSellWeight)

		stockMarketWatchData.Orders.SellAvgPricePercent = helper.Divide(
			stockMarketWatchData.Orders.SellAvgPrice-float32(stockMarketWatchData.StockPrices.PriceYesterday),
			float32(stockMarketWatchData.StockPrices.PriceYesterday)) * 100

		stockMarketWatchData.Orders.BuyAvgPricePercent = helper.Divide(
			stockMarketWatchData.Orders.BuyAvgPrice-float32(stockMarketWatchData.StockPrices.PriceYesterday),
			float32(stockMarketWatchData.StockPrices.PriceYesterday)) * 100

		marketWatchData[identifier] = stockMarketWatchData
	}
	MarketWatchData = marketWatchData
}

func convertStringToTile(stringTime string) time.Time {
	le := utf8.RuneCountInString(stringTime) - 1
	hour := cast.ToInt(stringTime[:le-3])
	minute := cast.ToInt(stringTime[le-3 : le-1])
	second := cast.ToInt(stringTime[le-1:])
	return time.Date(y, time.Month(m), d, hour, minute, second, 0, helper.GetIranTimeZone())
}

func FindStockInMarketWatch(identifier string) MarketWatchModel {
	stock, _ := MarketWatchData[identifier]
	return stock
}
