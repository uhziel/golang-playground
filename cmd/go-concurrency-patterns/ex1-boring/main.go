package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 1; i < math.MaxInt; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
	}
}

func main() {
	go boring("boring")

	fmt.Println("I'm listening.")
	time.Sleep(3 * time.Second)
	fmt.Println("I'm leaving. You are too boring.")
}
