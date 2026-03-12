package main

import (
	"fmt"
	"time"
)

func concurrentWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func WorkerPool(){
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i:= 1; i <=3; i=i+1 {
		go concurrentWorker(i, jobs, results)
	}

	for i:= range numJobs {
		jobs <- i + 1
	}

	close(jobs)

	for range numJobs {
		<-results
	}
}