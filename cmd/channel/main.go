package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 4, -5, 7, 10, 8, 3, 2}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	s1, s2 := <-c, <-c
	fmt.Println(s1, s2, s1+s2)
}
