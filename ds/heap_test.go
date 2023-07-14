package ds

import (
	"fmt"
	"testing"
)

func TestHeap_InsertMax(t *testing.T) {
	h := &Heap{}
	h.InsertMax(1)
	h.InsertMax(2)
	h.InsertMax(3)
	fmt.Println(h.Data)
}
