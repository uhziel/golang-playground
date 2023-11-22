package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(c)
	}
}

func echo(w io.Writer, text string, delay time.Duration) {
	fmt.Fprintln(w, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(w, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(w, "\t", strings.ToLower(text))
}

func handleConn(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		go echo(c, scanner.Text(), 1*time.Second)
	}
}
