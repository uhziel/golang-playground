package main

import (
	"fmt"
	"time"
)

func player(done <-chan bool, table chan *Ball, name string) {
	for {
		select {
		case <-done:
			return
		case ball, ok := <-table:
			if !ok {
				break
			}
			ball.hits++
			fmt.Printf("hit %d: %s\n", ball.hits, name)
			time.Sleep(time.Second)
			select {
			case <-done:
				return
			case table <- ball:
			}
		}
	}
}

type Ball struct {
	hits int
}

func main() {
	done := make(chan bool)
	table := make(chan *Ball)
	defer close(done)
	defer close(table)
	go player(done, table, "ping")
	go player(done, table, "pong")

	table <- &Ball{}
	time.Sleep(10 * time.Second)
	<-table
	panic("show me the stack")
}
