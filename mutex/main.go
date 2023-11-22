package main

import (
	"fmt"
	"sync"
)

var (
	mu  sync.Mutex
	num int = 101
)

func Lock() {
	mu.Lock()
}

func main() {
	mu.Lock()
	//Lock()
	fmt.Println(num)
	defer mu.Unlock()
}
