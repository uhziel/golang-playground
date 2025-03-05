package main

import (
	"fmt"
)

func main() {
	testforr(nil)
	testforr([]int{})
	testforr([]int{101})

	testlen(nil)
	testlen([]int{})
	testlen([]int{101})

	testcap(nil)
	testcap([]int{})
	testcap([]int{101})
}

func testforr(s []int) {
	fmt.Printf("## forr s: %#v\n", s)
	for _, v := range s {
		fmt.Println(v)
	}
}

func testlen(s []int) {
	fmt.Printf("## len() s: %#v\n", s)
	fmt.Println(len(s))
}

func testcap(s []int) {
	fmt.Printf("## cap() s: %#v\n", s)
	fmt.Println(cap(s))
}
