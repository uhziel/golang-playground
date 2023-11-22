package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

func getInstance() *Singleton {
	if instance == nil {
		//或者使用 init() 也行，只要它允许在第一次引用包这个比较早的时期
		once.Do(func() {
			instance = &Singleton{}
			fmt.Println("Creating single instance now.")
		})
	}

	return instance
}
