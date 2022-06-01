package main

import (
	"fmt"
	"sync"
)

type Stock struct {
	Symbol string
	Price  int
}

func (s *Stock) String() string {
	return fmt.Sprintf("%s: %d", s.Symbol, s.Price)
}

var stockPool = sync.Pool{
	New: func() interface{} {
		return &Stock{}
	},
}

func main() {
	var stocks []*Stock
	for i := 0; i < 10; i++ {
		stock := stockPool.Get().(*Stock)
		stock.Symbol = "GOOG"
		stock.Price = i
		stockPool.Put(stock)
		stocks = append(stocks, stock)
	}
	fmt.Printf("%v", stocks)
}
