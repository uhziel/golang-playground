package main

import (
	"fmt"
	"sync"
	"time"
)

const workerNum = 5

func main() {
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func(id int) {
			defer wg.Done()
			<-ch
			fmt.Printf("working workID=%d\n", id)
		}(i)
	}

	time.AfterFunc(3*time.Second, func() {
		close(ch)
	})

	wg.Wait()
	fmt.Println("process exited")
}
