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
		targetPrice := portfolio.Stocks[0].Price*(1.0-0.1*float64(len(portfolio.Stocks)))
		if targetPrice > market.Low {
			return true
		}
	}

	return false
}

func CanSell(market market.Market, portfolio portfolio.Portfolio) bool {
	if len(portfolio.Stocks) == 0 {
		return false
	}

	if portfolio.AveragePrice()*1.22 < market.High {
		return true
	}

	return false
}
