package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

type InnerConn struct {
	conn *websocket.Conn
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	ic := new(InnerConn)
	ic.conn = c

	defer c.Close()
	for {
		mt, message, err := ic.readMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ic.writeMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (c *InnerConn) readMessage() (messageType int, p []byte, err error) {
	mt, msg, err := c.conn.ReadMessage()
	if mt != websocket.BinaryMessage {
		panic("not binary message")
	}
	msg, err = decode(msg)
	if err != nil {
		panic(err)
	}
	return mt, msg, err
}

const (
	INTERESTING_HEADER uint32 = (0x1234ffff)
	ARM                uint16 = (0x2)
	YTUMMY             uint16 = (0x3)
	FOOTER             uint16 = (0x2)
)

func (c *InnerConn) writeMessage(messageType int, p []byte) error {
	if messageType != websocket.BinaryMessage {
		panic("not binary message")
	}
	msg := encode(p)
	return c.conn.WriteMessage(messageType, msg)
}

func encode(p []byte) []byte {
	negotiateLen := 10
	bodyBytes := make([]byte, len(p)+negotiateLen)
	binary.BigEndian.PutUint32(bodyBytes[:4], INTERESTING_HEADER)
	binary.BigEndian.PutUint16(bodyBytes[4:6], ARM)
	binary.BigEndian.PutUint16(bodyBytes[6:8], YTUMMY)
	binary.BigEndian.PutUint16(bodyBytes[8:10], FOOTER)
	copy(bodyBytes[10:], p)
	return p
}

func decode(p []byte) ([]byte, error) {
	ih := binary.BigEndian.Uint32(p[:4])
	arm := binary.BigEndian.Uint16(p[4:6])
	ytummy := binary.BigEndian.Uint16(p[6:8])
	footer := binary.BigEndian.Uint16(p[8:10])
	if ih != INTERESTING_HEADER || arm != ARM || ytummy != YTUMMY || footer != FOOTER {
		return nil, errors.New("not interesting header or arm or yummy or footer")
	}
	return p[10:], nil
}
