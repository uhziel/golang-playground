package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s say: %d", name, i)
			time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := boring("Lee")

	fmt.Println("I'm listening.")

	timeout := time.After(10 * time.Second)

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("You say too much.")
			return
		}
	}

	fmt.Println("I'm leaving. You are too boring.")
}
