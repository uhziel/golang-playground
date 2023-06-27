package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	mu       sync.Mutex
)

func getInstance() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
			fmt.Println("Creating single instance now.")
		}
	}

	return instance
}
