package main

import (
	"fmt"
)

func appendInt(slice []int, elems ...int) []int {
	n := len(slice)
	m := n + len(elems)
	if m > cap(slice) {
		tmp := make([]int, m*2)
		copy(tmp, slice)
		slice = tmp
	}
	slice = slice[0:m]
	copy(slice[n:m], elems)
	return slice
}

func main() {
	s := []int{1, 2, 3}
	s = appendInt(s, 4, 5)
	fmt.Println(s)
	s2 := appendInt(nil, s...)
	fmt.Println(s2)

	// func append(slice []Type, elems ...Type) []Type
	// 参数 elems 可以是 nil
	// 参数 slice 也可以是 nil
	s = append(s, nil...)
	fmt.Println(s)
	var nilSlice []int
	s3 := append(nilSlice, s...)
	fmt.Println(s3)
	s4 := append(nilSlice, nil...)
	fmt.Println(s4)
}
