package main

import (
	"context"
	"log"
	"time"
)

func main() {
	sb := &slidingblock{}
	sb.generateArray()
	defer sb.Print()

	b := &slidingblock{
		id:         0,
		reasonable: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	cpb1 := *b
	cpb2 := *b
	cpb3 := *b
	cpb4 := *b
	cpb5 := *b
	cpb5.reasonable = false
	sb.Pass(ctx, &cpb1)
	sb.Pass(ctx, &cpb2)
	sb.Pass(ctx, &cpb3)
	sb.Pass(ctx, &cpb4)
	sb.Pass(ctx, &cpb5)

	cpb6 := *b
	cpb7 := *b
	cpb8 := *b
	cpb9 := *b
	cpb10 := *b
	sb.Pass(ctx, &cpb6)
	sb.Pass(ctx, &cpb7)
	sb.Pass(ctx, &cpb8)
	sb.Pass(ctx, &cpb9)
	sb.Pass(ctx, &cpb10)
}

type Ruler interface {
	Pass(context.Context, *slidingblock)
}

type slidingblock struct {
	id           uint
	reasonable   bool
	sbs          []*slidingblock
	cheapPointer uint
}

// define: block reasonable is true
func (sb *slidingblock) Pass(ctx context.Context, b *slidingblock) {
	log.Printf("content:%v,addr:%p \n", b, &b)
	defer sb.explosd()
	if !b.reasonable {
		sb.push(b)
		// record current index
		sb.cheapPointer = uint(len(sb.sbs))
		return
	}
	sb.push(b)
}

func (sb *slidingblock) generateArray() {
	sbs := make([]*slidingblock, 0)
	sb.sbs = sbs
}

func (sb *slidingblock) push(b *slidingblock) {
	sb.sbs = append(sb.sbs, b)
}

// define: continuous 5 block is explosd
const blockExplosdLen = 5

func (sb *slidingblock) explosd() {
	if len(sb.sbs) < 5 {
		log.Println("not explosive enough")
		return
	}
	if len(sb.sbs)-int(sb.cheapPointer) >= blockExplosdLen {
		log.Println("explosive enough")
	}
}

func (sb *slidingblock) Print() {
	log.Printf("%+v", sb)
}

var _ Ruler = (*slidingblock)(nil)
