package main

import (
	"fmt"
	"sync/atomic"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	mgr := NewMgr()
	pipe := make(chan struct{}, 4)
	for _, v := range []int{1, 2, 3, 4} {
		work := NewWorker(mgr.get, mgr.set, pipe)
		switch v {
		case 1:
			go work.dosomthing("拖地")
		case 2:
			go work.dosomthing("洗衣")
		case 3:
			go work.dosomthing("做饭")
		case 4:
			go work.dosomthing("打杂")
		}
	}
	for {
		select {
		case <-pipe:
			for k, v := range mgr.cache.Items() {
				fmt.Printf("k:%s v:%v \n", k, v.Object)
			}
		default:
			// if len(pipe) <= 0 {
			// 	return
			// }
			time.Sleep(time.Microsecond * 10)
		}
	}
}

type Mgr struct {
	cache *cache.Cache
}

func NewMgr() *Mgr {
	return &Mgr{
		cache: cache.New(time.Second*5, time.Second*5),
	}
}

func (mgr *Mgr) get(k string) {
	mgr.cache.Get(k)
}

func (mgr *Mgr) set(k string, v string) {
	mgr.cache.SetDefault(k, v)
}

type Get func(k string)
type Set func(k, v string)

type Woker struct {
	id   int32
	get  Get
	set  Set
	pipe chan struct{}
}

var wid atomic.Int32

func NewWorker(get func(string), set func(string, string), pipe chan struct{}) *Woker {
	wid.Add(1)
	return &Woker{
		id:   wid.Load(),
		get:  get,
		set:  set,
		pipe: pipe,
	}
}

func tostr(src int32) string {
	return fmt.Sprintf("%d", src)
}

func (w *Woker) dosomthing(v string) {
	w.get(tostr(w.id))
	defer w.set(tostr(w.id), v)
	w.pipe <- struct{}{}

}
