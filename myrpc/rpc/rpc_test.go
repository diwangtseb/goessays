package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	gob.Register(User{})
	addr := ":1234"
	s := NewServer(addr)
	s.Register("queryUser", queryUser)
	go s.Run()
	time.Sleep(time.Second * 1)
	ok := make(chan struct{})
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
		}
		client := NewClient(conn)
		var query func(int) (User, error)
		client.callRPC("queryUser", &query)
		u, err := query(1)
		if err != nil {
			panic(err)
		}
		ok <- struct{}{}
		fmt.Printf("%v", u)
	}()
	<-ok
}
