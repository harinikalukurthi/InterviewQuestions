package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d is processing job %d\n", id, job)
		time.Sleep(5*time.Second)
        // Perform your work here.
    }
}

func main() {
    numWorkers := 3
    numJobs := 10
    jobs := make(chan int, numJobs)
    var wg sync.WaitGroup

    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, jobs, &wg)
    }

    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }
    close(jobs)

    wg.Wait()
}