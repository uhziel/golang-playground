package main

import "time"

const (
	pendingQueueLen = 10
)

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

func Subscribe(f Fetcher) Subscription {
	s := &sub{
		fetcher: f,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	go s.run()
	return s
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

type fetchResult struct {
	items []Item
	next  time.Time
	err   error
}

func (s *sub) run() {
	var next time.Time
	var err error
	var pending []Item
	seen := make(map[string]bool)
	var fetchResultCh chan fetchResult

	for {
		var delay time.Duration
		if now := time.Now(); next.After(now) {
			delay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if fetchResultCh == nil && len(pending) <= pendingQueueLen {
			startFetch = time.After(delay)
		}

		var firstItem Item
		var updates chan<- Item
		if len(pending) > 0 {
			firstItem = pending[0]
			updates = s.updates
		}

		select {
		case errCh := <-s.closing:
			close(s.updates)
			errCh <- err
			return
		case <-startFetch:
			fetchResultCh = make(chan fetchResult)
			// fetchResultCh = make(chan fetchResult, 1) 很容易忽视，可能导致下面这个 goroutine 泄漏。
			// 加入 ctx 而不是缓存加 1 可能是更直觉且不容易出错的方法。
			go func() {
				items, next, err := s.fetcher.Fetch()
				fetchResultCh <- fetchResult{
					items: items,
					next:  next,
					err:   err,
				}
			}()
		case fetchResult := <-fetchResultCh:
			fetchResultCh = nil
			var items []Item
			items, next, err = fetchResult.items, fetchResult.next, fetchResult.err

			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}

			for _, item := range items {
				if !seen[item.uuid] {
					seen[item.uuid] = true

					pending = append(pending, item)
				}
			}

		case updates <- firstItem:
			pending = pending[1:]
		}
	}
}

// 功能正常但是未优化的版本
func (s *sub) run1() {
	var next time.Time
	var err error

	var pending []Item

	for {
		var delay time.Duration
		if now := time.Now(); next.After(now) {
			delay = next.Sub(now)
		}
		startFetch := time.After(delay)

		var firstItem Item
		var updates chan<- Item
		if len(pending) > 0 {
			firstItem = pending[0]
			updates = s.updates
		}

		select {
		case errCh := <-s.closing:
			close(s.updates)
			errCh <- err
			return
		case <-startFetch:
			var items []Item
			items, next, err = s.fetcher.Fetch()
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			pending = append(pending, items...)
		case updates <- firstItem:
			pending = pending[1:]
		}
	}
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errCh := make(chan error)
	s.closing <- errCh
	return <-errCh
}
