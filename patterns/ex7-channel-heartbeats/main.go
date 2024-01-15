package main

import (
	"fmt"
	"time"
)

func doWork(
	done <-chan struct{},
	heartbeatInterval time.Duration,
) (<-chan struct{}, <-chan time.Time) {
	heartbeatCh := make(chan struct{}, 1)
	workCh := make(chan time.Time)
	go func() {
		defer close(heartbeatCh)
		defer close(workCh)

		tick := time.Tick(heartbeatInterval)
		workTick := time.Tick(2 * heartbeatInterval)

		sendHeartbeat := func() {
			select {
			case heartbeatCh <- struct{}{}:
			default:
			}
		}

		sendWork := func(v time.Time) {
			for {
				select {
				case <-done:
					return
				case <-tick:
					sendHeartbeat()
				case workCh <- v:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-tick:
				sendHeartbeat()
			case v := <-workTick:
				sendWork(v)
			}
		}
	}()

	return heartbeatCh, workCh
}

func main() {
	done := make(chan struct{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	heartbeatCh, workCh := doWork(done, time.Second)
	for {
		select {
		case _, ok := <-heartbeatCh:
			if !ok {
				return
			}
			fmt.Println("heartbeat")
		case v, ok := <-workCh:
			if !ok {
				return
			}
			fmt.Println(v.Second())
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
