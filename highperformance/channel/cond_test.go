package channel

import (
	"testing"
	"time"
)

func TestNewInnerChan(t *testing.T) {
	nic := NewInnerChan()
	go nic.Read()
	go nic.Read()
	go nic.Read()
	nic.Write()
	time.Sleep(time.Second * 3)
}
