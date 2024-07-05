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
	ans := []Result{}
	ans = append(ans, Web(query))
	ans = append(ans, Image(query))
	ans = append(ans, Video(query))
	return ans
}

func main() {
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
