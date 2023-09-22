package eventengine

type Msg map[HeaderInfo][]BodyInfo

type HeaderInfo struct {
	Token string
}

type BodyInfo struct {
	Key   string
	Value string
}

type Collector interface {
	RecvMsg(*Msg)
	GetMsg() <-chan *Msg
}

type defaultCollector struct {
	chMsg chan *Msg
}

// GetMsg implements Collector.
func (dc *defaultCollector) GetMsg() <-chan *Msg {
	return dc.chMsg
}

// RecvMsg implements Collector.
func (dc *defaultCollector) RecvMsg(msg *Msg) {
	dc.chMsg <- msg
}

var _ Collector = (*defaultCollector)(nil)
