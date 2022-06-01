package main

import (
	"sync"
)

type Book struct {
	ID   string `json:"id"`
	Nmae string `json:"name"`
}

var bookPool = sync.Pool{
	New: func() interface{} {
		return new(Book)
	},
}
