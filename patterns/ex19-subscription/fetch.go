package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	uuid    string
	channel string
	title   string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func NewFakeFetcher(addr string) Fetcher {
	return &fakeFetcher{addr: addr}
}

type fakeFetcher struct {
	addr  string
	items []Item
}

var AllowDuplicates bool

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(30)) * 100 * time.Millisecond)

	item := Item{
		channel: f.addr,
		title:   fmt.Sprintf("item%d", len(f.items)),
	}
	item.uuid = item.channel + "/" + item.title

	f.items = append(f.items, item)

	if AllowDuplicates {
		items = f.items
	} else {
		items = []Item{item}
	}

	return
}
