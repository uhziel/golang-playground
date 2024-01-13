package main

import (
	"fmt"
	"math/rand"
)

func generator(done <-chan struct{}, nums []int) <-chan int {
	stream := make(chan int)
	go func() {
		defer close(stream)

		for _, num := range nums {
			select {
			case <-done:
				return
			case stream <- num:
			}
		}
	}()
	return stream
}

func add(done <-chan struct{}, n int, inStream <-chan int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for v := range inStream {
			select {
			case <-done:
				return
			case outStream <- v + n:
			}
		}
	}()
	return outStream
}

func multiply(done <-chan struct{}, n int, inStream <-chan int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-inStream:
				if !ok {
					return
				}
				outStream <- v * n
			}
		}
	}()
	return outStream
}

func repeat(done <-chan struct{}, values ...int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case outStream <- v:
				}
			}
		}
	}()
	return outStream
}

func take(done <-chan struct{}, inStream <-chan int, n int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case outStream <- <-inStream: // 这里有两个 chan 操作
			}
		}
	}()
	return outStream
}

func repeatFn(done <-chan struct{}, fn func() int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)

		for {
			select {
			case <-done:
				return
			case outStream <- fn():
			}
		}
	}()

	return outStream
}

func main() {
	nums := []int{
		1,
		3,
		2,
		10,
	}

	done := make(chan struct{})
	defer close(done)

	inCh := generator(done, nums)
	for v := range multiply(done, 2, add(done, 1, inCh)) {
		fmt.Println(v)
	}

	fmt.Println("take(done, repeat(done, 10, 11), 5)")
	for v := range take(done, repeat(done, 10, 11), 5) {
		fmt.Println(v)
	}

	fmt.Println("random")
	randNum := func() int {
		return rand.Intn(10)
	}
	for v := range take(done, repeatFn(done, randNum), 5) {
		fmt.Println(v)
	}
}
