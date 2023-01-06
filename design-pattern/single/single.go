package single

import (
	"log"
	"sync"
)

type SingleTon struct{}

var lock = &sync.Mutex{}
var singleTon *SingleTon

func NewSingleTone() *SingleTon {
	if singleTon == nil {
		lock.TryLock()
		defer lock.Unlock()
		singleTon = &SingleTon{}
		log.Println("entry single is nil")
		return singleTon
	}
	log.Println("entry single is not nil")
	return singleTon
}
