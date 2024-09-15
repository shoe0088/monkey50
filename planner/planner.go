package planner

import (
	"monkey50/market"
	"monkey50/portfolio"
)

func CanBuy(market market.Market, portfolio portfolio.Portfolio) bool {
	if len(portfolio.Stocks) == 0 {
		if market.RSI < 34 {
			return true
		}
	} else {
		lastStock := portfolio.Stocks[len(portfolio.Stocks)-1]
		if lastStock.Price*0.9 > market.Close {
			return true
		}
	}

	return false
}

func CanSell(market market.Market, portfolio portfolio.Portfolio) bool {
	if len(portfolio.Stocks) == 0 {
		return false
	}

	if portfolio.AveragePrice()*1.5 < market.Close {
		return true
	}

	return false
}
