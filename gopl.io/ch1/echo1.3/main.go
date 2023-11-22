package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func testForLoop() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("forLoop elapsed:", elapsed)
}

func testStringsJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("strings.Join() elapsed:", elapsed)
}

func main() {
	testForLoop()
	testStringsJoin()
}
