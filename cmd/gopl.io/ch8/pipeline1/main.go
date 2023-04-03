package main

import "fmt"

func main() {
	naturals, squares := make(chan int), make(chan int)

	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	go func() {
        for {
            n := <-naturals
            squares <- n * n
        }
	}()

	for {
		fmt.Println(<-squares)
	}
}
