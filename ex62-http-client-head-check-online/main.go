package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Head("http://localhost:6000/sync-to-s3")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	io.Copy(os.Stdout, resp.Body)
}
