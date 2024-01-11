package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	err  error
	resp *http.Response
}

func checkStatus(done chan struct{}, urls []string) <-chan Result {
	resultStream := make(chan Result)
	go func() {
		defer close(resultStream)

		for _, url := range urls {
			resp, err := http.Get(url)
			select {
			case <-done:
				return
			case resultStream <- Result{err, resp}:
			}
		}
	}()

	return resultStream
}

func main() {
	urls := []string{
		"http://www.example.com",
		"http://www.zhulei-bad.com",
	}

	done := make(chan struct{})
	time.AfterFunc(2*time.Second, func() {
		close(done)
	})

	for result := range checkStatus(done, urls) {
		if result.err != nil {
			fmt.Println(result.err)
			continue
		}
		fmt.Println(result.resp.Status)
	}
}
