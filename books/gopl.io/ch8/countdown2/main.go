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

	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("countdown aborted!")
		return
	}
	fmt.Println("launch.")
}
