package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan string

var (
	enterings = make(chan client)
	leavings  = make(chan client)
	messages  = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case c := <-enterings:
			clients[c] = true
		case c := <-leavings:
			delete(clients, c)
			close(c)
		case msg := <-messages:
			for c := range clients {
				c <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	c := make(client)

	enterings <- c

	go clientWriter(conn, c)
	who := conn.RemoteAddr().String()
	c <- "you are " + who + "\n"
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		messages <- who + ": " + msg + "\n"
	}

	leavings <- c
	conn.Close()
}

func clientWriter(conn net.Conn, c <-chan string) {
	for msg := range c {
		fmt.Fprint(conn, msg)
	}
}
