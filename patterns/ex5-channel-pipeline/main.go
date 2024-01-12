package main

import "fmt"

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
}
