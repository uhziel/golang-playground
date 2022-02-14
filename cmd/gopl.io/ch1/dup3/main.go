package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := os.Args[1:]
	for _, fileName := range fileNames {
		data, err := os.ReadFile(fileName)
		if err != nil {
			log.Println("dup3:", err)
			continue
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}
