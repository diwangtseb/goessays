package session

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Session struct {
	conn net.Conn
	opts sessionOption
}
type sessionOption struct {
	len uint8
}

type SessionOption interface {
	apply(*sessionOption)
}
type funcSO func(*sessionOption)

// apply implements SessionOption
func (fso funcSO) apply(s *sessionOption) {
	fso(s)
}

func WithLens(lens uint8) funcSO {
	return func(so *sessionOption) {
		fmt.Println(so.len, lens)
		so.len = lens
	}
}

func NewSession(conn net.Conn, ops ...SessionOption) *Session {
	so := sessionOption{}
	for _, o := range ops {
		o.apply(&so)
	}
	return &Session{
		conn: conn,
		opts: so,
	}
}

// protocol: bit map [header|body] bit lens[1<<2:]
func (s *Session) Write(data []byte) error {
	m := make([]byte, int(s.opts.len)+len(data))
	binary.BigEndian.PutUint32(m[:s.opts.len], uint32(len(data)))
	copy(m[s.opts.len:], data)
	_, err := s.conn.Write(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) Read() ([]byte, error) {
	h := make([]byte, s.opts.len)
	_, err := s.conn.Read(h)
	if err != nil {
		panic(err)
	}
	lens := binary.BigEndian.Uint32(h)
	data := make([]byte, lens)
	_, err = s.conn.Read(data)
	if err != nil {
		panic(err)
	}
	return data, nil
}

// func caseRun() {
// 	startCh := make(chan struct{})
// 	defer close(startCh)
// 	go func() {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				fmt.Println(err)
// 			}
// 		}()

// 		fmt.Println("start write")
// 		startCh <- struct{}{}
// 		lis, err := net.Listen("tcp", ":1234")
// 		if err != nil {
// 			panic(err)
// 		}
// 		conn, err := lis.Accept()
// 		if err != nil {
// 			panic(err)
// 		}

// 		ses := NewSession(conn, WithLens(10))
// 		for {
// 			err = ses.Write([]byte("hello"))
// 			if err != nil {
// 				panic(err)
// 			}
// 			time.Sleep(time.Second * 1)
// 		}
// 	}()
// 	go func() {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				fmt.Println(err)
// 			}
// 		}()
// 		<-startCh
// 		fmt.Println("start read")
// 		conn, err := net.Dial("tcp", ":1234")
// 		if err != nil {
// 			panic(err)
// 		}
// 		ses := NewSession(conn, WithLens(10))
// 		for {
// 			b, err := ses.Read()
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(string(b))
// 			panic(err)
// 		}
// 	}()
// 	exit := make(chan os.Signal)
// 	signal.Notify(exit, os.Interrupt)
// 	select {
// 	case <-exit:
// 		fmt.Println("secure exit")
// 		os.Exit(1)
// 	}
// }

// func main() {
// 	go caseRun()
// 	time.Sleep(time.Second * 20)
// }
