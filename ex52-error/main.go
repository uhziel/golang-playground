package main

import "fmt"

func main() {
	fmt.Println(fmt.Errorf("open fail: %w", nil))
}
