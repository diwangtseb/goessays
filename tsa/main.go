package main

import (
	"fmt"
)

func main() {
	type Result struct {
		Score float64
	}
	result := []Result{
		{
			Score: 300000.2156577674,
		},
		{
			Score: 30000.2156577673,
		},
		{
			Score: 3000.215657742,
		},
	}
	score := decimalIntFromPower(result[0].Score)
	index := 0
	var powerValue float64
	for _, v := range result {
		powerValue += v.Score
		if index > 1 {
			val := decimalIntFromPower(v.Score)
			if val < score {
				score = val
			}
		}
		index++
	}
	fmt.Println(score, powerValue)
}

func decimalIntFromPower(power float64) float64 {
	dp := power - float64(int(power))
	return dp
}
