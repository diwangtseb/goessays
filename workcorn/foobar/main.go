package main

import (
	"fmt"
	hello "helloworld"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	hello.Hello()
	fmt.Println("end")
	now := time.Now().Local()
	fmt.Println(timeToCronFromat(now))
	c := cron.New(cron.WithSeconds())
	c.AddFunc(timeToCronFromat(now), func() { fmt.Println("regularly perform") })
	c.Start()
	time.Sleep(time.Second * 10)
}

func timeToCronFromat(t time.Time) string {
	second := t.Second()
	minute := t.Minute()
	hour := t.Hour()
	day := t.Day()
	formatCronStr := fmt.Sprintf("%d %d %d %d * *", second, minute, hour, day)
	return formatCronStr
}
