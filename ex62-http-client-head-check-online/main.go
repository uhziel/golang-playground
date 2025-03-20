package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	head("http://localhost:6000/sync-to-s3")
	tcpdial("localhost:6000")
}

func head(url string) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Head(url)
	if err != nil {
		fmt.Println("head: err", err)
		return
	}

	fmt.Println("head:", resp.Status)
	fmt.Println("head:", resp.Header)
	io.Copy(os.Stdout, resp.Body)
}

func tcpdial(addr string) {
	conn, err := net.DialTimeout("tcp", addr, 2 * time.Second)
	if err != nil {
		fmt.Println("tcpdial: err", err)
		return
	}
	defer conn.Close()
	fmt.Println("tcpdial: success")
}
