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
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func First(query string, searches ...Search) Result {
	ch := make(chan Result)
	for i := range searches {
		go func(i int) {
			ch <- searches[i](query)
		}(i)
	}

	return <-ch
}

func Google(query string) []Result {
	ch := make(chan Result)
	go func() {
		ch <- First(query, Web1, Web2)
	}()
	go func() {
		ch <- First(query, Image1, Image2)
	}()
	go func() {
		ch <- First(query, Video1, Video2)
	}()

	ans := []Result{}
	timeout := time.After(500 * time.Millisecond)
LOOP:
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			ans = append(ans, result)
		case <-timeout:
			break LOOP
		}
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
