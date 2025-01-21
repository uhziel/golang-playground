package main

import (
	"context"
	"io"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	client := &http.Client{} // 真实项目使用时可以使用 http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
