package main

import (
	"fmt"
	"math/rand"
)

func func1() (a int, b string) {
	a = 1
	b = "abc"

	r := rand.Intn(2)
	if r == 0 {
		return
	}

	return 0, "efg"
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(func1())
	}
}
