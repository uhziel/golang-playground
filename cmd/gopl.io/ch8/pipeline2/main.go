package main

import "fmt"

func main() {
	naturals, squares := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		/*
			for {
				n, ok := <-naturals
				if !ok {
					close(squares)
					break
				}
				squares <- n * n
			}
		*/
		for v := range naturals {
			squares <- v * v
		}
		close(squares)
	}()

	/*
		for {
			ans, ok := <-squares
			if !ok {
				break
			}
			fmt.Println(ans)
		}
	*/
	for v := range squares {
		fmt.Println(v)
	}
}
