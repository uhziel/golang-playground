package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func boring(msg string, output chan<- string) {
	for i := 1; i < math.MaxInt; i++ {
		output <- fmt.Sprintln(msg, i)
		time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
	}
}

func main() {
	ch := make(chan string)
	go boring("boring", ch)

	fmt.Println("I'm listening.")
	for i := 0; i < 3; i++ {
		fmt.Println("You say:", <-ch)
	}
	fmt.Println("I'm leaving. You are too boring.")
}
