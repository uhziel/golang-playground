package main

import (
	"fmt"
	"math/rand"
	"sync"
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

func fanIn(chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	chAllInOne := make(chan string)
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for v := range ch {
				chAllInOne <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(chAllInOne)
	}()
	return chAllInOne
}

func main() {
	ch := fanIn(boring("Lee"), boring("John"))

	fmt.Println("I'm listening.")

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("I'm leaving. You are too boring.")
}
