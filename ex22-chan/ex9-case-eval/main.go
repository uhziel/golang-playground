package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	time.AfterFunc(time.Second, func() {
		close(done)
	})

	ch := make(chan int)

	start := time.Now()
	select {
	case <-done:
	case ch <- longTime():
	}
	fmt.Println(time.Since(start))
}

func longTime() int {
	time.Sleep(3 * time.Second)
	return 1
}
