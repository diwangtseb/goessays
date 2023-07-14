package ds

import (
	"fmt"
	"testing"
)

func TestTree_CreateTreeByArray(t *testing.T) {
	arr := []int{5, 3, 2, 4, 7, 6, 8}
	r := InsertBbtByRecursionArray(arr)
	r2 := StackTraverse(r)
	r3 := InsertBstByRecursionArray(arr)
	fmt.Println("height->", r3.Height())
	r3.Dfs()
	fmt.Println("------")
	r3.Bfs()
	r4 := StackTraverse(r3)
	fmt.Println(r, r2, r3, r4)
}
