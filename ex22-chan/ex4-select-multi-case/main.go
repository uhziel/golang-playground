package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(time.Second)
	ticker := time.Tick(time.Second)
	closedCh := make(chan int)
	close(closedCh)
	for {
		select {
		case t := <-timer:
			fmt.Println("time.After", t)
		case v, ok := <-closedCh:
			fmt.Println("closed v:", v, "ok:", ok)
			time.Sleep(time.Second)
		case t := <-ticker:
			fmt.Println("time.Tick", t)
		}
	}
}
