package ds

import (
	"fmt"
	"testing"
)

func TestLink_Insert(t *testing.T) {
	l := Link{}
	l.Insert(1)
	l.Insert(3)
	l.Insert(2)
	l.Insert(4)
	fmt.Println(l)
	r := l.Traverse()
	fmt.Println(r)
}
