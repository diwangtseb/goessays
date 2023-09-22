package eventengine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_ee(t *testing.T) {
	buffer := 10
	collect := defaultCollector{
		chMsg: make(chan *Msg, buffer),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	go func(ctx context.Context) {
		bit := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("go done")
				return
			default:
				duration := time.Second * 1
				headerInfo := HeaderInfo{
					Token: fmt.Sprintf("%d", bit),
				}
				bodInfo := []BodyInfo{
					{
						Key:   "rush",
						Value: "b",
					},
					{
						Key:   "ts",
						Value: "now",
					},
				}
				msgOrigion := map[HeaderInfo][]BodyInfo{
					headerInfo: bodInfo,
				}
				collect.RecvMsg((*Msg)(&msgOrigion))
				bit++
				time.Sleep(duration)
			}

		}
	}(ctx)
	eengine := eventEngine{
		collector: &collect,
	}
	eengine.Handle(ctx)
}
