package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}
