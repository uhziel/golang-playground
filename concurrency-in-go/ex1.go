package main

import (
	"fmt"
	"strings"
)

func main() {
	var stdoutBuilder strings.Builder
	defer func() { fmt.Println(stdoutBuilder.String()) }()

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuilder, "producer done.")
		for i := 0; i < 5; i++ {
			intStream <- i
			fmt.Fprintf(&stdoutBuilder, "send %d\n", i)
		}
	}()

	for v := range intStream {
		fmt.Fprintln(&stdoutBuilder, v)
	}
	fmt.Fprintln(&stdoutBuilder, "cosumer done.")
}
