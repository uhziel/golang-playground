package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
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
			case outStream <- <-inStream:
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

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func filterPrime(done <-chan struct{}, inStream <-chan int) <-chan int {
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

				if isPrime(v) {
					outStream <- v
				}
			}
		}
	}()
	return outStream
}

func fanIn(done <-chan struct{}, inStreams ...<-chan int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		wg := sync.WaitGroup{}
		wg.Add(len(inStreams))
		for _, inStream := range inStreams {
			go func(inStream <-chan int) {
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					case outStream <- <-inStream:
					}
				}
			}(inStream)
		}
		wg.Wait()
	}()
	return outStream
}

func orDone(done <-chan struct{}, inStream <-chan int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
			case v, ok := <-inStream:
				if !ok {
					return
				}

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

func tee(done <-chan struct{}, inStream <-chan int) (_, _ <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		defer close(out1)
		defer close(out2)

		for v := range orDone(done, inStream) {
			o1, o2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case o1 <- v:
					o1 = nil
				case o2 <- v:
					o2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func bridge(done <-chan struct{}, inChCh <-chan <-chan int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for {
			var inCh <-chan int
			select {
			case <-done:
				return
			case _inCh, ok := <-inChCh:
				if !ok {
					return
				}
				inCh = _inCh
			}

			for v := range orDone(done, inCh) {
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

	fmt.Println("tee(done, take(done, repeat(done, 10, 11), 5))")
	out1, out2 := tee(done, take(done, repeat(done, 10, 11), 5))
	for v := range out1 {
		fmt.Println(v, <-out2)
	}

	fmt.Println("random prime")
	randNum := func() int {
		return rand.Intn(100000000)
	}
	start := time.Now()
	for v := range take(done, filterPrime(done, repeatFn(done, randNum)), 3) {
		fmt.Println(v)
	}
	fmt.Printf("random prime done. time: %v\n", time.Since(start))

	// fan-in
	fmt.Println("random prime after fan-out and fan-in")
	start = time.Now()
	randIntStream := repeatFn(done, randNum)
	numCPU := runtime.NumCPU()
	primeStreams := make([]<-chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		primeStreams[i] = filterPrime(done, randIntStream)
	}
	for v := range take(done, fanIn(done, primeStreams...), 3) {
		fmt.Println(v)
	}
	fmt.Printf("random prime done after fan-out and fan-in. time: %v\n", time.Since(start))

	// bridge
	genVals := func() <-chan <-chan int {
		outChCh := make(chan (<-chan int))
		go func() {
			defer close(outChCh)
			for i := 0; i < 3; i++ {
				outCh := make(chan int, 1)
				outCh <- i
				close(outCh)

				outChCh <- outCh
			}
		}()
		return outChCh
	}
	for v := range bridge(done, genVals()) {
		fmt.Println(v)
	}
}
