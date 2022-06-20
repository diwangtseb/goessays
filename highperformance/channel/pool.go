package channel

import "sync"

type User struct {
	Id   int
	Name string
}

func NewUserPool() *sync.Pool {
	pool := &sync.Pool{
		New: func() any {
			return &User{}
		},
	}
	return pool
}
