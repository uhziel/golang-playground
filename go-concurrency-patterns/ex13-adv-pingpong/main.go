package main

import (
	"fmt"
	"time"
)

type Ball struct {
	Hits int
}

func player(name string, ch chan *Ball) {
	for {
		t := <-ch
		t.Hits++
		fmt.Println(name, t.Hits)
		time.Sleep(200 * time.Millisecond)
		ch <- t
	}
}

func main() {
	ch := make(chan *Ball)
	go player("ping", ch)
	go player("pong", ch)

	ch <- &Ball{}
	time.Sleep(time.Second)
	<-ch

	panic("show me the stack.")
}
