package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	topic   string
	content string
}

type SubscribeInfo struct {
	topic string
	ch    chan<- Event
}

func generator(done <-chan struct{}) <-chan Event {
	ch := make(chan Event)
	go func() {
		defer close(ch)
		events := []Event{
			{
				topic:   "topic1",
				content: "1",
			},
			{
				topic:   "topic2",
				content: "abc",
			},
		}
		for _, e := range events {
			select {
			case <-done:
				return
			case ch <- e:
			}
		}
	}()
	return ch
}

type Demux struct {
	m             map[string]chan<- Event
	done          <-chan struct{}
	in            <-chan Event
	subscribeCh   chan SubscribeInfo
	unsubscribeCh chan string
}

func NewDemux(done <-chan struct{}, in <-chan Event) *Demux {
	d := &Demux{
		m:             make(map[string]chan<- Event),
		done:          done,
		in:            in,
		subscribeCh:   make(chan SubscribeInfo),
		unsubscribeCh: make(chan string),
	}

	go d.run()
	return d
}

func (d *Demux) Subscribe(topic string, ch chan<- Event) bool {
	select {
	case <-d.done:
		return false
	case d.subscribeCh <- SubscribeInfo{topic, ch}:
		return true
	}
}

func (d *Demux) Unsubscribe(topic string) bool {
	select {
	case <-d.done:
		return false
	case d.unsubscribeCh <- topic:
		return true
	}
}

func (d *Demux) run() {
	defer close(d.subscribeCh)
	defer close(d.unsubscribeCh)

	for {
		select {
		case <-d.done:
			return
		case info := <-d.subscribeCh:
			d.m[info.topic] = info.ch
		case topic := <-d.unsubscribeCh:
			delete(d.m, topic)
		case e := <-d.in:
			out, ok := d.m[e.topic]
			if !ok {
				continue
			}
			select {
			case <-d.done:
				return
			case out <- e:
			}
		}
	}
}

func main() {
	done := make(chan struct{})
	time.AfterFunc(3*time.Second, func() {
		close(done)
	})

	in := generator(done)

	demux := NewDemux(done, in)

	out1, out2 := make(chan Event), make(chan Event)
	demux.Subscribe("topic1", out1)
	demux.Subscribe("topic2", out2)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case v := <-out1:
				fmt.Println(v)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case v := <-out2:
				fmt.Println(v)
			}
		}
	}()

	wg.Wait()
	fmt.Println("finished.")
	panic("show me the stacks")
}
