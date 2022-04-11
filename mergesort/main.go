package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func GetNumbers(in ...int) <-chan int {
	out := make(chan int, cap(in))
	go func() {
		for _, n := range in {
			out <- n
		}
		close(out)
	}()
	return out
}

func MemInSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		mem := make([]int, 0, cap(in))
		for n := range in {
			mem = append(mem, n)
		}
		sort.Ints(mem)
		for _, n := range mem {
			out <- n
		}
		close(out)
	}()
	return out
}

func MergeSort(in ...<-chan int) []int {
	out := make([]int, 0, cap(in))
	for _, c := range in {
		for n := range c {
			out = append(out, n)
			fmt.Println(out)
		}
	}
	sort.Ints(out)
	return out
}

var count = 2

const rate = 2

var out chan int = make(chan int)

func mockylotto() {

}

func iswinlotto() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%100 < count
}

func Get(id int) <-chan int {
	if iswinlotto() {
		out <- id
	}
	return out
}

func main() {
	// p1 := MemInSort(GetNumbers(4, 2, 7, 5, 1, 9))
	// p2 := MemInSort(GetNumbers(7, 3, 6, 0, 8))
	// out := MergeSort(p1, p2)
	// fmt.Println("out ->>", out)
	// for _, v := range out {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Printf("\n")
	go func() {
		for v := range out {
			fmt.Println("out --->", v)
			count--
		}
	}()
	for i := 0; i < 100000; i++ {
		go Get(i)
	}
	time.Sleep(time.Second * 10)
}
