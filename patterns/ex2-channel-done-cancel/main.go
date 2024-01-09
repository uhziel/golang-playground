package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan struct{}, stringStream <-chan string) <-chan struct{} {
	terminated := make(chan struct{})
	go func() {
		defer close(terminated)

		for {
			select {
			case s := <-stringStream:
				fmt.Println(s)
			case <-done:
				return
			case <-time.After(time.Second):
				fmt.Println("tick")
			}
		}
	}()
	return terminated
}

func main() {
	start := time.Now()
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(5 * time.Second)
	}()

	terminated := doWork(done, nil)
	<-terminated
	fmt.Printf("finished at %v", time.Since(start))
}
