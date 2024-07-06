package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("exec after timeout")
		case <-done:
			fmt.Println("exec after done")
			return
			//default:
		}

		fmt.Println("exec non-preemptable code")
		time.Sleep(time.Second)
	}
}
