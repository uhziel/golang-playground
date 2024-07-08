package main

import (
	"fmt"
	"time"
)

func generator(done <-chan bool) <-chan int {
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

func tee(done <-chan bool, in <-chan int) (out1, out2 <-chan int) {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		defer close(ch1)
		defer close(ch2)

		for v := range in {
			ch1, ch2 := ch1, ch2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case ch1 <- v:
					ch1 = nil
				case ch2 <- v:
					ch2 = nil
				}
			}
		}
	}()

	out1 = ch1
	out2 = ch2
	return
}

func main() {
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		close(done)
	}()

	ch1, ch2 := tee(done, generator(done))
	for v := range ch1 {
		fmt.Println(v, <-ch2)
	}

	panic("show me the stack")
}
