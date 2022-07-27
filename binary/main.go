package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	b := make([]byte, 8)
	fmt.Printf("% x\n", b)
	binary.BigEndian.PutUint16(b[0:], 0x03e8)
	binary.BigEndian.PutUint32(b[2:], 0x07d0)
	fmt.Printf("% x\n", b)
}
