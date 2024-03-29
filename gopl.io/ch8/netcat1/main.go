package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Print(err)
	}
}
