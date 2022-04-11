package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FluCard struct {
	name    string //flucard's name
	rate    int    // flucard's rate
	count   int    // flucard's count
	totle   int    //flucard's totle
	minCode int    // flucard's minCode
	maxCode int    // flucard's maxCode
}

const CODERANGE = 1999

type FluCards []FluCard

var FLUCARDS = FluCards{
	FluCard{name: "富强福卡", totle: 100, rate: 20, count: 100, minCode: 0, maxCode: 99},
	FluCard{name: "民主福卡", totle: 200, rate: 20, count: 200, minCode: 100, maxCode: 299},
	FluCard{name: "爱国福卡", totle: 300, rate: 20, count: 300, minCode: 300, maxCode: 599},
	FluCard{name: "敬业福卡", totle: 400, rate: 20, count: 400, minCode: 600, maxCode: 999},
	FluCard{name: "友善福卡", totle: 1000, rate: 20, count: 1000, minCode: 1000, maxCode: 1999},
}

func NewFluCards() *FluCards {
	return &FLUCARDS
}

func (fcs *FluCards) SendFluCard(code int, wg *sync.WaitGroup) {
	for _, fc := range *fcs {
		if code >= fc.minCode && code <= fc.maxCode {
			if IsWinPrice(code, fc.rate) {
				fc.count--
				fmt.Println("恭喜你获得了", fc.name, "福卡", fc.totle, fc.count)
			}
		}
	}
	wg.Done()
}

func IsWinPrice(code, rate int) bool {
	return rand.Intn(CODERANGE)/100 <= rate
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	fcs := NewFluCards()
	for i := 0; i < 90; i++ {
		wg.Add(1)
		r := rand.Intn(CODERANGE)
		go fcs.SendFluCard(r, &wg)
	}
	wg.Wait()
}
