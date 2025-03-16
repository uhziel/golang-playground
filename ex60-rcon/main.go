package main

import (
	"fmt"
	"log"

	"github.com/gorcon/rcon"
)

func main() {
	conn, err := rcon.Dial("192.168.31.64:25575", "123456")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	exec(conn, "help")
	exec(conn, "help 2")
	exec(conn, "help2")
}

func exec(conn *rcon.Conn, cmd string) {
	response, err := conn.Execute(cmd)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(response)	
}
