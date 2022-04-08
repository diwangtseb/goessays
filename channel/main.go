package main

import "time"

//cycle
func CreateTicker() *time.Ticker {
	return time.NewTicker(time.Second)
}

//once
func CreateTimer() *time.Timer {
	return time.NewTimer(time.Second)
}

func main() {
	ticker := CreateTicker()
	timer := CreateTimer()
	for {
		select {
		case <-ticker.C:
			println("tick")
		case <-timer.C:
			println("timer")
		}
	}
}
