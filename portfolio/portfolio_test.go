package portfolio

import (
	"testing"
	"time"
)

func TestPortfolio_averagePrice(t *testing.T) {
	// テストケース1: 株がない場合
	t.Run("NoStocks", func(t *testing.T) {
		p := Portfolio{
			Budget: 10000,
			Cash:   10000,
			Stocks: []Stock{},
		}

		got := p.AveragePrice()
		want := 0.0

		if got != want {
			t.Errorf("averagePrice() = %v, want %v", got, want)
		}
	})

	// テストケース2: 1つの株式のみの場合
	t.Run("OneStock", func(t *testing.T) {
		p := Portfolio{
			Budget: 10000,
			Cash:   9000,
			Stocks: []Stock{
				{
					Date:   time.Now(),
					Price:  100,
					Amount: 10,
				},
			},
		}

		got := p.AveragePrice()
		want := 100.0

		if got != want {
			t.Errorf("averagePrice() = %v, want %v", got, want)
		}
	})

	// テストケース3: 複数の株式がある場合
	t.Run("MultipleStocks", func(t *testing.T) {
		p := Portfolio{
			Budget: 10000,
			Cash:   8000,
			Stocks: []Stock{
				{
					Date:   time.Now(),
					Price:  100,
					Amount: 10,
				},
				{
					Date:   time.Now(),
					Price:  200,
					Amount: 5,
				},
			},
		}

		got := p.AveragePrice()
		want := (100*10 + 200*5) / float64(10+5)

		if got != want {
			t.Errorf("averagePrice() = %v, want %v", got, want)
		}
	})
}
