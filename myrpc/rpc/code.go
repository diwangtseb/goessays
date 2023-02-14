package rpc

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct {
	Name string
	Args []interface{}
}

func encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	bufEncode := gob.NewEncoder(&buf)
	if err := bufEncode.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	bufDecode := gob.NewDecoder(buf)
	var data RPCData
	if err := bufDecode.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
