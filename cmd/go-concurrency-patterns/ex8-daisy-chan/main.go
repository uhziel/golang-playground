package main

import "fmt"

func f(left chan<- int, right <-chan int) {
	v := 1 + <-right
	left <- v
}

func main() {
	leftmost := make(chan int)
	left, right := leftmost, leftmost
	for i := 0; i < 100; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func() { right <- 1 }()
	fmt.Println(<-leftmost)
}
