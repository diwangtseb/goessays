package main

import (
	"encoding/json"
	"testing"
)

var buf, _ = json.Marshal(Book{
	ID:   "1",
	Nmae: "2",
})

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := bookPool.Get().(*Book)
		json.Unmarshal(buf, stu)
		bookPool.Put(stu)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Book{}
		json.Unmarshal(buf, stu)
	}
}
