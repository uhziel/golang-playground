package main

import "fmt"

func forRangeSlice(s []string) {
	const id = "forRangeSlice"
	for range s {
		fmt.Println(id)
	}

	for i := range s {
		fmt.Println(id, i)
	}

	for i, v := range s {
		fmt.Println(id, i, v)
	}

	fmt.Println("")
}

func forRangeMap(m map[string]int) {
	const id = "forRangeMap"
	for range m {
		fmt.Println(id)
	}

	for k := range m {
		fmt.Println(id, k)
	}

	for k, v := range m {
		fmt.Println(id, k, v)
	}

	fmt.Println("")
}

func forRangeString(s string) {
	const id = "forRangeString"

	for range s {
		fmt.Println(id)
	}

	for i := range s {
		fmt.Println(id, i)
	}

	for i, v := range s {
		fmt.Printf("%s %d %c codepoint(%X)\n", id, i, v, v)
	}

	fmt.Println("")
}

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	return ch
}

func forRangeChannel() {
	const id = "forRangeChannel"

	for range generator() {
		fmt.Println(id)
	}

	for v := range generator() {
		fmt.Println(id, v)
	}
}

func main() {
	forRangeSlice([]string{"a", "b", "c"})
	forRangeMap(map[string]int{
		"a": 11,
		"b": 12,
		"c": 13,
	})
	forRangeString("中国")
	forRangeChannel()
}
