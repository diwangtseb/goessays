package websocket

import (
	"bufio"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"net/http"
)

const (
	// 是否是最后一个数据帧
	finalBit = 1 << 7
	// 是否需要进行掩码处理
	maskBit = 1 << 7

	// 文本数据帧类型
	TextMessage = 1
	// 关闭数据帧类型
	CloseMessage = 8
)

// WebSocket 链接
type Conn struct {
	writeBuf []byte
	maskKey  [4]byte

	conn net.Conn
}

func newConn(conn net.Conn) *Conn {
	return &Conn{conn: conn}
}

func (c *Conn) Close() {
	c.conn.Close()
}

// 发送数据，只支持发送文本数据，且不支持分片
func (c *Conn) SendData(data []byte) {
	length := len(data)
	c.writeBuf = make([]byte, 10+length)

	// 数据开始和结束的位置
	payloadStart := 2

	// 数据帧的第一个字节, 不支持分片，且值能发送文本类型数据
	// 所以二进制位为 %b1000 0001
	// b0 := []byte{0x81}
	c.writeBuf[0] = byte(TextMessage) | finalBit

	// 数据帧第二个字节，服务器发送的数据不需要进行掩码处理
	switch {
	case length >= 65536:
		c.writeBuf[1] = byte(0x00) | 127
		binary.BigEndian.PutUint64(c.writeBuf[payloadStart:], uint64(length))
		// 需要 8 byte 来存储数据长度
		payloadStart += 8
	case length > 125:
		c.writeBuf[1] = byte(0x00) | 126
		binary.BigEndian.PutUint16(c.writeBuf[payloadStart:], uint16(length))
		// 需要 2 byte 来存储数据长度
		payloadStart += 2
	default:
		c.writeBuf[1] = byte(0x00) | byte(length)
	}
	copy(c.writeBuf[payloadStart:], data[:])
	c.conn.Write(c.writeBuf[:payloadStart+length])
}

// 读取数据
func (c *Conn) ReadData() (data []byte, err error) {

	var b [8]byte
	// 读取数据帧的前两个字节
	if _, err := c.conn.Read(b[:2]); err != nil {
		return nil, err
	}

	// 开始解析第一个字节, 是否还有后续数据帧
	final := b[0]&finalBit != 0
	// 不支持数据分片
	if !final {
		log.Println("Recived fragmented frame, not support")
		return nil, errors.New("not support fragmented message")
	}

	// 数据帧类型
	frameType := int(b[0] & 0xf)
	// 如果关闭类型，则关闭链接
	if frameType == CloseMessage {
		c.conn.Close()
		log.Println("Recived closed message, connection will be closed")
		return nil, errors.New("recived closed message")
	}
	if frameType != TextMessage {
		return nil, errors.New("only support text message")
	}
	// 检查数据帧是否被掩码处理
	mask := b[1]&maskBit != 0

	// 数据长度
	payloadLen := int64(b[1] & 0x7F)
	dataLen := int64(payloadLen)
	// 根据payload length 判断数据的真实长度
	switch payloadLen {
	case 126:
		if _, err := c.conn.Read(b[:2]); err != nil {
			return nil, err
		}
		dataLen = int64(binary.BigEndian.Uint16(b[:2]))
	case 127:
		if _, err := c.conn.Read(b[:8]); err != nil {
			return nil, err
		}
		dataLen = int64(binary.BigEndian.Uint64(b[:8]))
	}

	log.Printf("Read data length: %d, payload length %d", payloadLen, dataLen)
	// 读取 mask key
	if mask {
		if _, err := c.conn.Read(c.maskKey[:]); err != nil {
			return nil, err
		}
	}

	// 读取数据内容
	p := make([]byte, dataLen)
	if _, err := c.conn.Read(p); err != nil {
		return nil, err
	}
	if mask {
		maskBytes(c.maskKey, 0, p)
	}
	return p, nil
}

// 将 http 链接升级到 websocket 链接
func Upgrade(w http.ResponseWriter, r *http.Request) (c *Conn, err error) {
	// 是否是 GET 方法
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil, errors.New("websocket: method not GET")
	}
	// 检查 Sec-WebSocket-Version 版本
	if values := r.Header["Sec-Websocket-Version"]; len(values) == 0 || values[0] != "13" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, errors.New("websocket: version != 13")
	}

	// 检查 Connection 和 Upgrade
	if !tokenListContainsValue(r.Header, "Connection", "upgrade") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, errors.New("websocket: could not find connection header with token 'upgrade'")
	}
	if !tokenListContainsValue(r.Header, "Upgrade", "websocket") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, errors.New("websocket: could not find connection header with token 'websocket'")
	}

	// 计算 Sec-WebSocket-Accpet 的值
	challengeKey := r.Header.Get("Sec-Websocket-Key")
	if challengeKey == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, errors.New("websocket: key missing or blank")
	}

	var (
		netConn net.Conn
		br      *bufio.Reader
	)

	h, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, errors.New("websocket: response dose not implement http.Hijacker")
	}
	var rw *bufio.ReadWriter
	netConn, rw, err = h.Hijack()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}
	br = rw.Reader

	if br.Buffered() > 0 {
		netConn.Close()
		return nil, errors.New("websocket: client sent data before handshake is complete")
	}

	// 构造握手成功后返回的 response
	p := []byte{}
	p = append(p, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: "...)
	p = append(p, computeAcceptKey(challengeKey)...)
	p = append(p, "\r\n\r\n"...)

	if _, err = netConn.Write(p); err != nil {
		netConn.Close()
		return nil, err
	}
	log.Println("Upgrade http to websocket successfully")
	conn := newConn(netConn)
	return conn, nil
}
