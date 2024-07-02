package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
https://gobyexample.com/reading-files
https://golang.org/pkg/bufio/#example_Scanner_lines
这种使用scanner.Scan()的方式感觉是最合适的。支持大文件读取，也方便后续的处理
*/

func main() {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatalf("Open /proc/cpuinfo fail. %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
