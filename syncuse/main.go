package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// 互斥锁
var mu sync.Mutex

// 记录红包已抢总金额
var sum uint

// 当前有效红包列表，int64是红包唯一ID，[]uint是红包里面随机分到的金额（单位分）
var packageList *sync.Map = new(sync.Map)

// 并发任务量
const TaskNum = 5

type task struct {
	id       uint32
	callback chan uint
}

// 构造任务队列
var ChanTasks []chan task = make([]chan task, TaskNum)

// 分配的随机数
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 为了方便查看，我们使用打印日志文档的方式
var logger *log.Logger

func main() {
	// 初始化日志
	initLog()
	// 发红包
	id := 1
	// 发红包的单位设置为元
	money, num := 12345, 10000
	SetRedPack(id, money, num)
	// 启动监听任务，通道里有消息时可以及时执行
	for i := 0; i < TaskNum; i++ {
		ChanTasks[i] = make(chan task)
	}
	go GetPackageMoney(ChanTasks)
	begin := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < num+5; i++ {
		wg.Add(1)
		go GetRedPack(id, &wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("抢到的总金额为%d", sum)
	fmt.Println("用时", end.Sub(begin))

}

// 初始化日志
func initLog() {
	f, _ := os.Create("./lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func SetRedPack(id, money, num int) {
	// 乘 100 是因为我们用整数表示金额 1 就代表 0.01 元
	moneyTotal := int(money * 100)
	// 金额分配算法
	leftMoney := moneyTotal
	leftNum := num

	list := make([]uint, num)
	// 大循环开始，只要还有没分配的名额，继续分配
	for leftNum > 0 {
		if leftNum == 1 {
			// 最后一个名额，把剩余的全部给它
			list[num-1] = uint(leftMoney)
			break
		}
		// 剩下的最多只能分配到1分钱时，不用再随机
		if leftMoney == leftNum {
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		// 每次对剩余金额的1%-55%随机，最小1，最大就是剩余金额55%（需要给剩余的名额留下1分钱的生存空间）
		rMoney := int(2 * float64(leftMoney) / float64(leftNum))
		m := r.Intn(rMoney) + 1
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}
	//packageList[id] = list
	packageList.Store(uint32(id), list)
}

func GetRedPack(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 并发安全字典取 value
	list1, ok := packageList.Load(uint32(id))
	list := list1.([]uint)
	// 没取到就代表红包不存在
	if !ok || len(list) < 1 {
		logger.Printf("红包不存在,id=%d\n", id)
	}
	// 构造抢红包任务
	callback := make(chan uint)
	t := task{
		id:       uint32(id),
		callback: callback,
	}
	// 我们之前启动了 16 个 goroutine 分别监听不同的 channel ，也是为了一个分流的作用
	ChanTasks[sum%TaskNum] <- t
	money := <-t.callback
	if money <= 0 {
		logger.Printf("很遗憾，你没有抢到红包\n")
	} else {
		logger.Printf("恭喜你抢到一个红包，金额为:%d\n", money)
		mu.Lock()
		defer mu.Unlock()
		sum += money
	}
}

func GetPackageMoney(ChanTask []chan task) {
	for {
		select {
		case t := <-ChanTask[0]:
			GetMoney(t)
		case t := <-ChanTask[1]:
			GetMoney(t)
		case t := <-ChanTask[2]:
			GetMoney(t)
		case t := <-ChanTask[3]:
			GetMoney(t)
		case t := <-ChanTask[4]:
			GetMoney(t)
		default:
			continue
		}
	}
}

func GetMoney(t task) {
	id := t.id
	l, ok := packageList.Load(id)
	if ok && l != nil {
		list := l.([]uint)
		// 从红包金额中随机得到一个
		i := r.Intn(len(list))
		money := list[i]
		// 更新红包列表中的信息
		if len(list) > 1 {
			if i == len(list)-1 {
				packageList.Store(uint32(id), list[:i])
			} else if i == 0 {
				packageList.Store(uint32(id), list[1:])
			} else {
				packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
			}
		} else {
			packageList.Delete(uint32(id))
		}
		t.callback <- money
		//return fmt.Sprintf("恭喜你抢到一个红包，金额为:%d\n", money)
	} else {
		t.callback <- 0
	}
}
