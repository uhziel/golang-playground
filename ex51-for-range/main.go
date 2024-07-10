package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	for range s {
		fmt.Println("hello")
	}

	for i := range s {
		fmt.Println("hello", i)
	}

	for i, v := range s {
		fmt.Println("hello", i, v)
	}
}
