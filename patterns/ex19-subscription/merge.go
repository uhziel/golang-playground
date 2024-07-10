package main

import "sync"

func Merge(subs ...Subscription) Subscription {
	m := &merge{
		updates:    make(chan Item),
		subs:       subs,
		notifyDone: make(chan struct{}),
		err:        make(chan error),
	}
	go m.run()
	return m
}

type merge struct {
	updates    chan Item
	subs       []Subscription
	notifyDone chan struct{}
	err        chan error
}

func (m *merge) run() {
	var wg sync.WaitGroup

	wg.Add(len(m.subs))
	for _, sub := range m.subs {
		go func(sub Subscription) {
			defer wg.Done()
			for {
				var item Item
				select {
				case item = <-sub.Updates():
				case <-m.notifyDone:
					m.err <- sub.Close()
					return
				}

				select {
				case m.updates <- item:
				case <-m.notifyDone:
					m.err <- sub.Close()
					return
				}
			}
		}(sub)
	}

	wg.Wait()
	close(m.updates)
}

func (m *merge) Updates() <-chan Item {
	return m.updates
}

func (m *merge) Close() error {
	close(m.notifyDone)

	var firstErr error
	for range m.subs {
		if err := <-m.err; err != nil {
			if firstErr == nil {
				firstErr = err
			}
		}
	}

	return firstErr
}
