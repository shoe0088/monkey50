package planner

import (
	"testing"
	"time"

	"monkey50/market"
	"monkey50/portfolio"
)

func TestCanBuy(t *testing.T) {
	// テストケース1: 持株がない場合でRSIが34未満
	t.Run("NoStocks_RSI_Under_34", func(t *testing.T) {
		marketData := market.Market{
			RSI:   30,
			Close: 100,
		}
		port := portfolio.Portfolio{
			Stocks: []portfolio.Stock{},
		}

		if !CanBuy(marketData, port) {
			t.Errorf("Expected to be able to buy, but got false")
		}
	})

	// テストケース2: 持株がないがRSIが34以上
	t.Run("NoStocks_RSI_Above_34", func(t *testing.T) {
		marketData := market.Market{
			RSI:   40,
			Close: 100,
		}
		port := portfolio.Portfolio{
			Stocks: []portfolio.Stock{},
		}

		if CanBuy(marketData, port) {
			t.Errorf("Expected not to be able to buy, but got true")
		}
	})

	// テストケース3: 持株があり、最後の株価が市場価格の90%未満
	t.Run("HasStocks_LowerThan90Percent", func(t *testing.T) {
		marketData := market.Market{
			Close: 91,
		}
		port := portfolio.Portfolio{
			Stocks: []portfolio.Stock{
				{
					Date:   time.Now(),
					Price:  100,
					Amount: 10,
				},
			},
		}

		if !CanBuy(marketData, port) {
			t.Errorf("Expected to be able to buy, but got false")
		}
	})

	// テストケース4: 持株があり、最後の株価が市場価格の90%以上
	t.Run("HasStocks_HigherThan90Percent", func(t *testing.T) {
		marketData := market.Market{
			Close: 89,
		}
		port := portfolio.Portfolio{
			Stocks: []portfolio.Stock{
				{
					Date:   time.Now(),
					Price:  100,
					Amount: 10,
				},
			},
		}

		if CanBuy(marketData, port) {
			t.Errorf("Expected not to be able to buy, but got true")
		}
	})
}

func TestCanSell(t *testing.T) {
	// テストケース1: 持株がない場合
	t.Run("NoStocks", func(t *testing.T) {
		marketData := market.Market{
			Close: 100,
		}
		port := portfolio.Portfolio{
			Stocks: []portfolio.Stock{},
		}

		if CanSell(marketData, port) {
			t.Errorf("Expected not to be able to sell, but got true")
		}
	})

	// // テストケース2: 持株があり、平均取得価格の130%以上で市場価格が売却可能
	// t.Run("HasStocks_CanSellAt130Percent", func(t *testing.T) {
	// 	marketData := market.Market{
	// 		Close: 130,
	// 	}
	// 	port := portfolio.Portfolio{
	// 		Stocks: []portfolio.Stock{
	// 			{
	// 				Date:   time.Now(),
	// 				Price:  100,
	// 				Amount: 10,
	// 			},
	// 		},
	// 	}

	// 	// モックとして平均価格を100と仮定
	// 	port.AveragePrice = func() float64 {
	// 		return 100
	// 	}

	// 	if !canSell(marketData, port) {
	// 		t.Errorf("Expected to be able to sell, but got false")
	// 	}
	// })

	// // テストケース3: 持株があり、市場価格が平均取得価格の130%未満で売却不可
	// t.Run("HasStocks_CannotSellAtLessThan130Percent", func(t *testing.T) {
	// 	marketData := market.Market{
	// 		Close: 120,
	// 	}
	// 	port := portfolio.Portfolio{
	// 		Stocks: []portfolio.Stock{
	// 			{
	// 				Date:   time.Now(),
	// 				Price:  100,
	// 				Amount: 10,
	// 			},
	// 		},
	// 	}

	// 	// モックとして平均価格を100と仮定
	// 	port.AveragePrice = func() float64 {
	// 		return 100
	// 	}

	// 	if canSell(marketData, port) {
	// 		t.Errorf("Expected not to be able to sell, but got true")
	// 	}
	// })
}
