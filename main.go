package main

import (
	"monkey50/planner"
	"monkey50/portfolio"
	"monkey50/printer"
)

func main() {
	filePath := "soxl_2014.csv"
	p := portfolio.Portfolio{
		InitialBudget: 220000,
		Budget: 220000,
		Cash:   220000,
	}
	markets, _ := printer.ImportCSV(filePath)
	p.Start = markets[0].Date

	for _, m := range markets {
		if planner.CanBuy(m, p) {
			p.Buy(m)
		}
		if planner.CanSell(m, p) {
			p.Sell(m)
		}
	}

	p.Report(markets[len(markets)-1])
}