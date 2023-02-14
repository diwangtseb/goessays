package rpc

import (
	ms "myrpc/session"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		session := ms.NewSession(c.conn, ms.WithLens(10))
		reqRpc := RPCData{
			Name: rpcName,
			Args: inArgs,
		}
		b, err := encode(reqRpc)
		if err != nil {
			panic(err)
		}
		err = session.Write(b)
		if err != nil {
			panic(err)
		}
		respBytes, err := session.Read()
		if err != nil {
			panic(err)
		}
		respRpc, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		outArgs := make([]reflect.Value, 0, len(respRpc.Args))
		for i, arg := range respRpc.Args {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
