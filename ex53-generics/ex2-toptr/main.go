package main

import "fmt"

func ToPtr[T any](v T) *T {
	return &v
}

func main() {
	v := ToPtr(3)
	fmt.Printf("%T %v\n", v, v)
}
