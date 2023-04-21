package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("%s say: %d", name, i)
			time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
		}
		close(ch)
	}()
	return ch
}

func main() {
	chLee := boring("Lee")
	chJohn := boring("John")

	fmt.Println("I'm listening.")

	for msg := range chLee {
		fmt.Println(msg)
	}
	for msg := range chJohn {
		fmt.Println(msg)
	}

	/* 方法1
	for i := 0; i < 2; i++ {
		fmt.Println(<-chLee)
		fmt.Println(<-chJohn)
	}
	*/
	fmt.Println("I'm leaving. You are too boring.")
}
