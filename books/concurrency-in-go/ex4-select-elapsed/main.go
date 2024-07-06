package main

import (
	"fmt"
	"time"
)

func main() {
	var ch <-chan bool
	start := time.Now()
	select {
	case <-ch:
	default:
		fmt.Printf("select default elapsed=%v\n", time.Since(start))
	}
}
