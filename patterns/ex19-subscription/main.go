package main

import (
	"fmt"
	"time"
)

func main() {
	sub := Merge(
		Subscribe(NewFakeFetcher("google.com")),
		Subscribe(NewFakeFetcher("bing.com")),
	)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("start to close")
		fmt.Println("closed", sub.Close())
	})

	for item := range sub.Updates() {
		fmt.Println(item.channel, item.title)
	}

	panic("show me the stacks")
}
