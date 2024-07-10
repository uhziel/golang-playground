/*
 * 和 context pkg 像，但是相反
 */
package main

import (
	"fmt"
	"time"
)

/* 会出现重复 close(ch) 而导致的 panic
func or2(channels ...<-chan any) <-chan any {
	out := make(chan any)

	for _, ch := range channels {
		go func(ch <-chan any) {
			<-ch
			close(out)
		}(ch)
	}

	return out
}
*/

func or(channels ...<-chan any) <-chan any {
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
			case <-or(append(channels[3:], orCh)...):
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
	for v := range or(doneAfter1Sec, doneAfter3Sec) {
		fmt.Println(v)
	}
	fmt.Printf("done after %v\n", time.Since(start))
}
