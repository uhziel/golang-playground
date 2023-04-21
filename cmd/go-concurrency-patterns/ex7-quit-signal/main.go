package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s say: %d", name, i):
			case <-quit:
				fmt.Println("clean up")
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
		}
		close(ch)
	}()
	return ch
}

func main() {
	quit := make(chan string)
	ch := boring("Lee", quit)

	fmt.Println("I'm listening.")

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

	quit <- "Bye"
	fmt.Println("Lee say:", <-quit)
}
