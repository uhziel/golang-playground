package main

import (
	"fmt"
	"time"
)

func main() {
	var ch <-chan bool
	start := time.Now()
	select {
	case <-ch:
	case <-time.After(3 * time.Second):
		fmt.Printf("timer fire time=%v\n", time.Since(start))
	}
}
