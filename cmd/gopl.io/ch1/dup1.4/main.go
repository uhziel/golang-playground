package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lineLocations := make(map[string][]string)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		countLines("-", os.Stdin, lineLocations)
	} else {
		for _, fileName := range fileNames {
			file, err := os.Open(fileName)
			if err != nil {
				log.Printf("dup1.4: %v\n", err)
				continue
			}
			countLines(fileName, file, lineLocations)
			file.Close()
		}
	}

	for line, locations := range lineLocations {
		if len(locations) > 1 {
			fmt.Printf("%s\t%v\n", line, locations)
		}
	}
}

func countLines(fileName string, file *os.File, lineLocations map[string][]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		lineLocations[line] = append(lineLocations[line], fileName)
	}
}
