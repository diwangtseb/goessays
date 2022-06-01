package main

import (
	"sync"
)

type Book struct {
	ID     string `json:"id"`
	Nmae   string `json:"name"`
	Page   int    `json:"page"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

var bookPool = sync.Pool{
	New: func() interface{} {
		return new(Book)
	},
}
