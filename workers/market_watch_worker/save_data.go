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

	for stockId, data := range marketWatchData {
		// save stock price
		prices = append(prices, models.StockPrices{
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
		})

		//save orders
		saveOrders(&orders, data, cast.ToUint(stockId))

		// save stock info
		if stockIsNew(cast.ToUint(stockId), stocksInDb) {
			newStocks = append(newStocks, models.Stock{
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

	}
	//todo update stock info
	if len(newStocks) != 0 {
		models.Stock{}.DB().Create(&newStocks)
	}
	models.StockPrices{}.DB().Create(&prices)
	models.OrderItems{}.DB().Create(&orders)
}

func stockIsNew(stockId uint, stocksInDb []models.Stock) bool {
	for _, stock := range stocksInDb {
		if stock.ID == stockId {
			return false
		}
	}
	return true
}

func saveOrders(orders *[]models.OrderItems, data market_watch.MarketWatchModel, stockId uint) {
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
