package eventengine

import (
	"context"
	"fmt"
)

type EventHandler interface {
	Handle(ctx context.Context)
}

type eventEngine struct {
	collector Collector
}

// Handle implements EventHandler.
func (ee *eventEngine) Handle(ctx context.Context) {
	msg := ee.collector.GetMsg()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handle done")
			return
		case m := <-msg:
			for header, body := range *m {
				if header.Token == "" {
					continue
				}
				fmt.Println(header, body)
			}
		}
	}

}

var _ EventHandler = (*eventEngine)(nil)
