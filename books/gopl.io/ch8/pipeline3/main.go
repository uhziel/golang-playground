package main

import "fmt"

func main() {
	naturals, squares := make(chan int), make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
}

func squarer(out <-chan int, in chan<- int) {
	for v := range out {
		in <- v * v
	}
	close(in)
}

func printer(out <-chan int) {
	for v := range out {
		fmt.Println(v)
	}
}

// 我是按照 chan 的视角看 in/out 的
// gopl 是按照操作者的身份看 in/out
//
//     func counter(out chan int)
//     func squarer(out, in chan int)
//     func printer(in chan int)
