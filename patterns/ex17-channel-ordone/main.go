package main

import (
	"fmt"
	"time"
)

func generator(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		for i := 0; ; i++ {
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

func ordone(done <-chan struct{}, in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-done:
				case ch <- v:
				}
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

	ch := generator(done)

	for v := range ordone(done, ch) {
		fmt.Println(v)
	}
}
