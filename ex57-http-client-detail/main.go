package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
