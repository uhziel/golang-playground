package main

import (
	"fmt"
)

const n = 100000

func main() {
	var ch1, ch2, ch3 = make(chan bool), make(chan bool), make(chan bool)
	var counter1, counter2, counter3 int
	close(ch1)
	close(ch2)
	close(ch3)

	for i := 0; i < n; i++ {
		select {
		case <-ch1:
			counter1++
		case <-ch2:
			counter2++
		case <-ch3:
			counter3++
		}
	}

	fmt.Println("counter1=", counter1)
	fmt.Println("counter2=", counter2)
	fmt.Println("counter3=", counter3)
}
