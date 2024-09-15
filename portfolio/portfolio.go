package portfolio

import (
	"fmt"
	"monkey50/market"
	"time"
)

type Portfolio struct {
	Budget        float64    // 予算 CAD
	Cash         float64    // 残高 CAD
	Stocks [] Stock  // 持株情報
}

type Stock struct {
	Date  time.Time  // 購入日
	Price float64  // 取得単価
	Amount int  // 購入数
}

func (p *Portfolio) Buy(market market.Market) {
	amount := p.Budget/5/market.Close
	total := market.Close*amount
	if p.Cash < total {
		return
	}

	stock := Stock{
		Date: market.Date,
		Price: market.Close,
		Amount: int(amount),
	}
	p.Stocks = append(p.Stocks, stock)
	p.Cash = p.Cash - total
	p.Report()
}

func (p *Portfolio) Sell(market market.Market) {
	if len(p.Stocks) == 0 {
		return
	}

	totalAmount := 0
	for _, s := range p.Stocks {
		totalAmount += s.Amount
	}

	profit := market.Close*float64(totalAmount)
	p.Cash += profit
	p.Stocks = []Stock{}
	p.Report()
}

func (p *Portfolio) AveragePrice() float64 {
	var totalPrice float64
	var totalAmount int

	for _, stock := range p.Stocks {
		totalPrice += stock.Price * float64(stock.Amount)
		totalAmount += stock.Amount
	}

	if totalAmount == 0 {
		return 0 // 株数が0の場合は平均価格を0とする
	}

	return totalPrice / float64(totalAmount)
}

// Report はPortfolioの内容を標準出力に出力する関数です
func (p Portfolio) Report() {
	var totalPrice float64
	for _, stock := range p.Stocks {
		totalPrice += stock.Price * float64(stock.Amount)
	}
	totalPrice += p.Cash

	fmt.Println("==== Portfolio Report ====")
	fmt.Printf("予算 (Budget): %.2f CAD\n", p.Budget)
	fmt.Printf("残高 (Cash): %.2f CAD\n", p.Cash)

	if len(p.Stocks) == 0 {
		fmt.Println("持株がありません。")
	} else {
		fmt.Println("持株情報:")
		for i, stock := range p.Stocks {
			fmt.Printf("  株 #%d\n", i+1)
			fmt.Printf("    購入日: %s\n", stock.Date.Format("2006-01-02"))
			fmt.Printf("    取得単価: %.2f CAD\n", stock.Price)
			fmt.Printf("    購入数: %d\n", stock.Amount)
		}
		fmt.Printf("    平均取得単価: %.2f\n", p.AveragePrice())
	}

	fmt.Printf("総資産: %.2f ％ \n", totalPrice)
	fmt.Printf("上昇率: %.2f ％ \n", totalPrice/p.Budget*100)
	fmt.Println("==========================")
}
