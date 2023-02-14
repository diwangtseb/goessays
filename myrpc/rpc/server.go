package rpc

import (
	"fmt"
	ms "myrpc/session"
	"net"
	"reflect"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

func (s *Server) Register(name string, f interface{}) {
	if _, ok := s.funcs[name]; ok {
		return
	}
	s.funcs[name] = reflect.ValueOf(f)
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		session := ms.NewSession(conn, ms.WithLens(10))
		b, err := session.Read()
		if err != nil {
			panic(err)
		}
		rpcData, err := decode(b)
		if err != nil {
			return
		}
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			return
		}
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		out := f.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		respRpcData := RPCData{
			Name: rpcData.Name,
			Args: outArgs,
		}
		bytes, err := encode(respRpcData)
		if err != nil {
			return
		}
		err = session.Write(bytes)
		if err != nil {
			panic(err)
		}
	}
}

type User struct {
	Name string
	Age  int
}

func queryUser(uid int) (User, error) {
	switch uid {
	case 1:
		return User{"user1", 22}, nil
	default:
		return User{
			Name: "",
			Age:  0,
		}, fmt.Errorf("uid:%d", uid)

	}
}
