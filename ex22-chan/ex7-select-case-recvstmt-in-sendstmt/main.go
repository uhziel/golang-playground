package main

import (
	"fmt"
	"time"
)

const (
	n = 100
)

func generator(done <-chan struct{}, n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func forward(done <-chan struct{}, in <-chan int) <-chan int {
	chTick := time.Tick(time.Nanosecond)
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- <-in:
			case <-chTick:
			}
		}

	}()
	return ch
}

func main() {
	done := make(chan struct{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	ch := forward(done, generator(done, n))

	for v := range ch {
		fmt.Println(v)
	}
}
