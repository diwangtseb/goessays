package main

import (
	"fmt"
	"sort"
)

func QuicklySort(in []int) []int {
	if len(in) < 2 {
		return in
	}
	first := in[0]
	low := make([]int, 0)
	hight := make([]int, 0)
	mid := make([]int, 0)
	for i := 1; i < len(in); i++ {
		if in[i] < first {
			low = append(low, in[i])
		} else if in[i] > first {
			hight = append(hight, in[i])
		} else {
			mid = append(mid, in[i])
		}
	}
	low, hight = QuicklySort(low), QuicklySort(hight)
	return append(append(low, first), append(mid, hight...)...)
}

func BubbleSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in)-1-i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
	return in
}

func main() {
	// quicklysort
	//arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(QuicklySort(arr))
	//// bubblesort
	//arrbubble := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(BubbleSort(arrbubble))
	someMap := map[string]int{
		"0": 10,
		"1": 9,
		"2": 8,
		"3": 1,
		"4": 1,
	}
	keys := []string{}
	for k := range someMap {
		keys = append(keys, k)
	}
	slice := make([]struct {
		Key   string
		Count int
	}, 0)
	sort.Strings(keys)
	for _, v := range keys {
		slice = append(slice, struct {
			Key   string
			Count int
		}{
			Key:   v,
			Count: someMap[v],
		})
	}
	for k, v := range slice {
		fmt.Println(k, v)
	}

}
