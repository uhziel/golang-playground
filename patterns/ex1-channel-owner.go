package main

import "fmt"

func owner() <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()
	return intStream
}

func main() {
	for v := range owner() {
		fmt.Println(v)
	}
}
