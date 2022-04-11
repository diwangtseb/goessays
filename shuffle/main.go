package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// shuffleAlg(nums)
	// fmt.Printf("%v", nums)
	SendRedPkg(100, 5)
}

func shuffleAlg(nums []int) []int {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		j := i + rand.Intn(len(nums)-i)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func RandomRedPkg(money int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(money)
}

func SendRedPkg(money, count int) {
	sum := 0
	for i := 0; i < count; i++ {
		if i == count-1 {
			sum += money
			fmt.Println(sum)
			return
		}
		m := RandomRedPkg(money)
		money -= m
		sum += m
	}
}
