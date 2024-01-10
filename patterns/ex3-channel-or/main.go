package main

import (
	"fmt"
	"time"
)

func orDone(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orCh := make(chan any)
	go func() {
		defer close(orCh)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-orDone(append(channels[3:], orCh)...):
			}
		}
	}()

	return orCh
}

func newAfterCh(d time.Duration) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		time.Sleep(d)
	}()

	return ch
}

func main() {
	start := time.Now()

	doneAfter1Sec := newAfterCh(time.Second)
	doneAfter3Sec := newAfterCh(3 * time.Second)
	for v := range orDone(doneAfter1Sec, doneAfter3Sec) {
		fmt.Println(v)
	}
	fmt.Printf("done after %vs", time.Since(start))
}
