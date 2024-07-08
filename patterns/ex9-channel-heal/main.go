package main

import (
	"log"
	"time"
)

func or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orCh := make(chan struct{})
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

type startGoroutineFn func(done <-chan struct{}, heartbeatInterval time.Duration) <-chan struct{}

func newSteward(timeout time.Duration, workFn startGoroutineFn) startGoroutineFn {
	return func(done <-chan struct{}, heartbeatInterval time.Duration) <-chan struct{} {
		heartbeat := make(chan struct{}, 1)
		go func() {
			defer close(heartbeat)
			tick := time.Tick(heartbeatInterval)

			var workHeartbeatCh <-chan struct{}
			var workDone chan struct{}
			startWork := func() {
				workDone = make(chan struct{})
				workHeartbeatCh = workFn(or(done, workDone), heartbeatInterval)
			}
			startWork()

		outerLoop:
			for {
				timeoutCh := time.After(timeout)
				for {
					select {
					case <-done:
						return
					case <-tick:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-workHeartbeatCh:
						continue outerLoop
					case <-timeoutCh:
						log.Println("restart")
						close(workDone)
						startWork()
						continue outerLoop
					}
				}
			}
		}()

		return heartbeat
	}
}

func doWork(done <-chan struct{}, _ time.Duration) <-chan struct{} {
	log.Println("doWork starting")
	go func() {
		<-done
		log.Println("doWork done")
	}()
	return nil
}

func main() {
	done := make(chan struct{})
	time.AfterFunc(15*time.Second, func() { close(done) })

	steward := newSteward(3*time.Second, doWork)
	for range steward(done, 1*time.Second) {
		log.Println("steward tick")
	}
	log.Println("main done")
}
