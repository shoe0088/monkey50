package main

import (
	"fmt"
	"monkey50/planner"
	"monkey50/portfolio"
	"monkey50/printer"
)

func main() {
	filePath := "soxl.csv"
	p := portfolio.Portfolio{
		InitialBudget: 200000,
		Budget: 200000,
		Cash:   200000,
	}
	markets, _ := printer.ImportCSV(filePath)

	for _, m := range markets {
		fmt.Printf("%s \n", m.Date)
		if planner.CanBuy(m, p) {
			p.Buy(m)
		}
		if planner.CanSell(m, p) {
			p.Sell(m)
		}
	}

	p.Start = markets[0].Date
	p.End = markets[len(markets)-1].Date
	p.Report(markets[len(markets)-1])
}