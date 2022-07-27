package hello

import (
	"fmt"

	"github.com/robfig/cron"
)

func Hello() {
	c := cron.New()
	c.AddFunc("*/20 * * * *", func() { fmt.Println("every minutes") })
	c.Start()
	// c.Stop() // Stop the scheduler (does not stop any jobs already running).
	fmt.Println("Hello, world!")
}
