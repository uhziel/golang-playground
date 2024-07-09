package main

import "fmt"

func c1(i int) {
	fmt.Println(&i)
	f := func() func() {
		g := func() {
			fmt.Println(&i)
			fmt.Println(i) // 看到的还是最外层的 i
		}
		return g
	}

	f()()
}

func main() {
	c1(1)
	c1(2)
}
