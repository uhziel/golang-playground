package main

import (
	"fmt"
	"time"
)

func heartbeat() {
	tick := time.Tick(1 * time.Second)
	//for {
	//	fmt.Println("heartbeat")
	//	<-tick
	//}
	for {
		select {
		case <-tick:
			fmt.Println("heartbeat")
		}
	}
}

func main() {
	go heartbeat()
	select {}
}
