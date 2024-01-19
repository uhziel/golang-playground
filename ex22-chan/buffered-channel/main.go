package main

import "fmt"

// via http://127.0.0.1:3999/concurrency/3

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	// 放开下面行会因为 ch 已满而 block
	// ch <- 3
	fmt.Println(<-ch)
}
