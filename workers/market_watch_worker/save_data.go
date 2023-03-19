package market_watch_worker

import (
	"github.com/spf13/cast"
	"stock/models"
	"stock/stock_info_tse/market_watch"
)

func SaveMarketWatchData(noInput ...interface{}) {
	marketWatchData := market_watch.GetMarketWatch()
	var prices []models.StockPrices
	var orders []models.OrderItems
	var newStocks []models.Stock
	var stocksInDb []models.Stock
	models.Stock{}.DB().Find(&stocksInDb)
	var pricesIdDB = models.StockPrices{}.LastPriceOfStock()

	for stockId, data := range marketWatchData {
		// save stock info
		addStock(&newStocks, stocksInDb, data, cast.ToUint(stockId))

		// save stock price
		addPrices(&prices, &pricesIdDB, data, cast.ToUint(stockId))

		//save orders
		addOrders(&orders, data, cast.ToUint(stockId))

	}
	if len(newStocks) != 0 {
		models.Stock{}.DB().Create(&newStocks)
	}
	if len(prices) != 0 {
		models.StockPrices{}.DB().Create(&prices)
	}
	if len(orders) != 0 {
		models.OrderItems{}.DB().Create(&orders)
	}

}

func addPrices(prices *[]models.StockPrices, pricesIdDB *[]models.StockPrices,
	data market_watch.MarketWatchModel, stockId uint) {
	isNew := true
	for _, lastPrice := range *pricesIdDB {
		if lastPrice.StockId == stockId {
			if !data.StockPrices.Time.After(lastPrice.TransactionAt) {
				isNew = false
			}
			break
		}
	}
	if !isNew {
		return
	}
	priceModel := models.StockPrices{
		StockId:           cast.ToUint(stockId),
		PriceMin:          data.StockPrices.PriceMin,
		PriceMax:          data.StockPrices.PriceMax,
		PriceYesterday:    data.StockPrices.PriceYesterday,
		PriceFirst:        data.StockPrices.PriceFirst,
		PriceClose:        data.StockPrices.PriceClose,
		PriceLast:         data.StockPrices.PriceLast,
		TransactionCount:  data.StockPrices.TransactionCount,
		TransactionVolume: data.StockPrices.TransactionVolume,
		TransactionValue:  data.StockPrices.TransactionValue,
		TransactionAt:     data.StockPrices.Time,
	}
	*prices = append(*prices, priceModel)
	*pricesIdDB = append(*pricesIdDB, priceModel)
}

func addStock(stocks *[]models.Stock, stocksInDb []models.Stock, data market_watch.MarketWatchModel, stockId uint) {
	isNew := true
	for _, stock := range stocksInDb {
		if stock.ID == stockId {
			isNew = false
			break
		}
	}
	if !isNew {
		return
	}
	*stocks = append(*stocks, models.Stock{
		ID:                cast.ToUint(stockId),
		NameFa:            data.StockInfo.NameFa,
		SymbolDigitCode12: data.StockInfo.SymbolDigitCode12,
		SymbolName:        data.StockInfo.SymbolName,
		IndustryGroupCode: cast.ToInt(data.StockInfo.IndustryGroupCode),
		TotalStockNumber:  data.StockInfo.TotalStockNumber,
		MonthAverage:      data.StockInfo.MonthAverage,
		EPS:               data.StockInfo.EPS,
		PE:                data.StockInfo.PE,
		BaseVolume:        data.StockInfo.BaseVolume,
	})
}

func addOrders(orders *[]models.OrderItems, data market_watch.MarketWatchModel, stockId uint) {
	*orders = append(*orders, models.OrderItems{
		StockId:              stockId,
		RowNumber:            1,
		BNumber:              data.Orders.R1.Buy.Number,
		BPrice:               data.Orders.R1.Buy.Price,
		BVolume:              data.Orders.R1.Buy.Volume,
		BPricePercent:        data.Orders.R1.Buy.PricePercent,
		BIsInAuthorizedRange: data.Orders.R1.Buy.IsInAuthorizedRange,
		SNumber:              data.Orders.R1.Sell.Number,
		SPrice:               data.Orders.R1.Sell.Price,
		SVolume:              data.Orders.R1.Sell.Volume,
		SPricePercent:        data.Orders.R1.Sell.PricePercent,
		SIsInAuthorizedRange: data.Orders.R1.Sell.IsInAuthorizedRange,
		CreatedAt:            data.StockPrices.Time,
	})
	*orders = append(*orders, models.OrderItems{
		StockId:              stockId,
		RowNumber:            2,
		BNumber:              data.Orders.R2.Buy.Number,
		BPrice:               data.Orders.R2.Buy.Price,
		BVolume:              data.Orders.R2.Buy.Volume,
		BPricePercent:        data.Orders.R2.Buy.PricePercent,
		BIsInAuthorizedRange: data.Orders.R2.Buy.IsInAuthorizedRange,
		SNumber:              data.Orders.R2.Sell.Number,
		SPrice:               data.Orders.R2.Sell.Price,
		SVolume:              data.Orders.R2.Sell.Volume,
		SPricePercent:        data.Orders.R2.Sell.PricePercent,
		SIsInAuthorizedRange: data.Orders.R2.Sell.IsInAuthorizedRange,
		CreatedAt:            data.StockPrices.Time,
	})
	*orders = append(*orders, models.OrderItems{
		StockId:              stockId,
		RowNumber:            3,
		BNumber:              data.Orders.R3.Buy.Number,
		BPrice:               data.Orders.R3.Buy.Price,
		BVolume:              data.Orders.R3.Buy.Volume,
		BPricePercent:        data.Orders.R3.Buy.PricePercent,
		BIsInAuthorizedRange: data.Orders.R3.Buy.IsInAuthorizedRange,
		SNumber:              data.Orders.R3.Sell.Number,
		SPrice:               data.Orders.R3.Sell.Price,
		SVolume:              data.Orders.R3.Sell.Volume,
		SPricePercent:        data.Orders.R3.Sell.PricePercent,
		SIsInAuthorizedRange: data.Orders.R3.Sell.IsInAuthorizedRange,
		CreatedAt:            data.StockPrices.Time,
	})
	*orders = append(*orders, models.OrderItems{
		StockId:              stockId,
		RowNumber:            4,
		BNumber:              data.Orders.R4.Buy.Number,
		BPrice:               data.Orders.R4.Buy.Price,
		BVolume:              data.Orders.R4.Buy.Volume,
		BPricePercent:        data.Orders.R4.Buy.PricePercent,
		BIsInAuthorizedRange: data.Orders.R4.Buy.IsInAuthorizedRange,
		SNumber:              data.Orders.R4.Sell.Number,
		SPrice:               data.Orders.R4.Sell.Price,
		SVolume:              data.Orders.R4.Sell.Volume,
		SPricePercent:        data.Orders.R4.Sell.PricePercent,
		SIsInAuthorizedRange: data.Orders.R4.Sell.IsInAuthorizedRange,
		CreatedAt:            data.StockPrices.Time,
	})
	*orders = append(*orders, models.OrderItems{
		StockId:              stockId,
		RowNumber:            5,
		BNumber:              data.Orders.R5.Buy.Number,
		BPrice:               data.Orders.R5.Buy.Price,
		BVolume:              data.Orders.R5.Buy.Volume,
		BPricePercent:        data.Orders.R5.Buy.PricePercent,
		BIsInAuthorizedRange: data.Orders.R5.Buy.IsInAuthorizedRange,
		SNumber:              data.Orders.R5.Sell.Number,
		SPrice:               data.Orders.R5.Sell.Price,
		SVolume:              data.Orders.R5.Sell.Volume,
		SPricePercent:        data.Orders.R5.Sell.PricePercent,
		SIsInAuthorizedRange: data.Orders.R5.Sell.IsInAuthorizedRange,
		CreatedAt:            data.StockPrices.Time,
	})

}
