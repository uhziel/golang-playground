package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("started. worker:%d job:%d\n", id, job)
		started := time.Now()
		time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
		results <- job * job
		fmt.Printf("finished. worker:%d job:%d elapsed:%v\n", id, job, time.Since(started))
	}
}

func workerEfficient(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for job := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Printf("started. worker:%d job:%d\n", id, job)
			started := time.Now()
			time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
			results <- job * job
			fmt.Printf("finished. worker:%d job:%d elapsed:%v\n", id, job, time.Since(started))
		}(job)
	}

	wg.Wait()
}

func main() {
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for id := 0; id < 3; id++ {
		go workerEfficient(id, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}
