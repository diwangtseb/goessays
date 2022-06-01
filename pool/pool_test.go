package main

import (
	"encoding/json"
	"sync"
	"testing"
)

var buf, _ = json.Marshal(Book{
	ID:     "1",
	Nmae:   "2",
	Page:   0,
	Author: "",
	Price:  0,
})

var mapPool = sync.Pool{
	New: func() interface{} {
		return new(map[string]string)
	},
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := mapPool.Get().(*map[string]string)
		json.Unmarshal(buf, m)
		mapPool.Put(m)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var m map[string]string
		json.Unmarshal(buf, &m)
	}
}
