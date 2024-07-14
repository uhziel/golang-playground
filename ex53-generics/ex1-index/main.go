package main

import "fmt"

func Index[T comparable](s []T, v T) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{1, 2, 3}
	fmt.Println(Index(si, 3))
	fmt.Println(Index(si, 10))

	ss := []string{"a", "b", "c"}
	fmt.Println(Index(ss, "b"))
	fmt.Println(Index(ss, "bb"))
}
