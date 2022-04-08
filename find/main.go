package main

import "fmt"

func BinSerach(in []int, target int) int {
	low := 0
	high := len(in) - 1
	for low <= high {
		mid := (low + high) / 2
		if in[mid] > target {
			high = mid - 1
		} else if in[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 1024<<10)
	for i := 0; i < 1024<<10; i++ {
		arr[i] = i + 1
	}
	id := BinSerach(arr, 12)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}
