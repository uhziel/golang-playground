package main

import "fmt"

func filbonacci(n int) int {
	if n <= 1 {
		return n
	}

	return n * filbonacci(n-1)
}

func main() {
	fmt.Println(filbonacci(4))
}
