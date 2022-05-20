package hello

import (
	"fmt"

	"github.com/robfig/cron"
)

func Hello() {
	c := cron.New()
	c.AddFunc("* * * * *", func() { fmt.Println("every minutes") })
	c.Start()
	// c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
