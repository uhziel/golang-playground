package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Countdown. Press return to abort.")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
		case <-abort:
			fmt.Println("countdown aborted!")
			return
		}
	}

	fmt.Println("launch.")
}
