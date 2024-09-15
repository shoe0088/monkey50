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
		firstStock := portfolio.Stocks[0]
		if firstStock.Price*(1.0-0.1*float64(len(portfolio.Stocks))) > market.Close {
			return true
		}
	}

	return false
}

func CanSell(market market.Market, portfolio portfolio.Portfolio) bool {
	if len(portfolio.Stocks) == 0 {
		return false
	}

	if portfolio.AveragePrice()*1.22 < market.Close {
		return true
	}

	return false
}
