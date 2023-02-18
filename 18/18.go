package main

import (
	"fmt"
	"time"
)

type Counter struct {
	Count int
}

func (c *Counter) Counter() {
	for {
		c.Count++
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	var c Counter
	timeout := time.After(5 * time.Second)
	go c.Counter()
	for {
		select {
		case <-timeout:
			fmt.Println(c.Count)
			return

		}
	}
}
