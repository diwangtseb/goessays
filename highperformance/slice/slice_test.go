package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_Slice_Len_Cap(t *testing.T) {
	sliceInt := Slice()
	fmt.Printf("temp len is %d,cap is %d\n", len(sliceInt), cap(sliceInt))
	for i := 0; i < 10; i++ {
		sliceInt = append(sliceInt, 6)
		r := unsafe.Sizeof(sliceInt)
		fmt.Printf("sliceInt len is %d,cap is %d,sizeOf is %d\n ", len(sliceInt), cap(sliceInt), r)
	}
}

func Test_Slice_Assignment(t *testing.T) {
	sliceInt := Slice()
	sliceIntAnother := sliceInt

	// append origin
	sliceInt = append(sliceInt, 1)
	fmt.Printf("sliceInt len is %d,cap is %d\n", len(sliceInt), cap(sliceInt))
	fmt.Printf("sliceIntAnother len is %d,cap is %d\n", len(sliceIntAnother), cap(sliceIntAnother))
}

func Test_Slice_Release(t *testing.T) {
	sliceInt := Slice()
	sliceIntAnother := sliceInt
	sliceInt = nil
	fmt.Printf("sliceInt len is %d,cap is %d\n", len(sliceInt), cap(sliceInt))
	fmt.Printf("sliceIntAnother len is %d,cap is %d\n", len(sliceIntAnother), cap(sliceIntAnother))
}
