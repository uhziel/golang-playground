package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(200)*10) * time.Millisecond)
		return Result(fmt.Sprintf("(%s in %s)", query, kind))
	}
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Google(query string) []Result {
	ch := make(chan Result)
	go func() {
		ch <- Web(query)
	}()
	go func() {
		ch <- Image(query)
	}()
	go func() {
		ch <- Video(query)
	}()

	ans := []Result{}
	for i := 0; i < 3; i++ {
		ans = append(ans, <-ch)
	}
	return ans
}

func main() {
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
