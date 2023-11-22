package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk(ctx context.Context, wait time.Duration, msg string) {
	select {
	case <-time.After(wait):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	sleepAndTalk(ctx, 3*time.Second, "hello")
}
