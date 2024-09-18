package main

import (
	"fmt"
	"sync"
)

// Job represents the task to be processed (a slice of 10 integers)
type Job struct {
	data []int
}

// Worker represents a worker that processes jobs
type Worker struct {
	id         int
	jobChannel chan Job
	wg         *sync.WaitGroup
}

// NewWorker creates a new Worker
func NewWorker(id int, jobChannel chan Job, wg *sync.WaitGroup) Worker {
	return Worker{
		id:         id,
		jobChannel: jobChannel,
		wg:         wg,
	}
}

// Start begins the worker's job processing loop
func (w Worker) Start() {
	go func() {
		for job := range w.jobChannel {
			fmt.Printf("Worker %d processing job: %v\n", w.id, job.data)
			for _, num := range job.data {
				fmt.Printf("Worker %d prints: %d\n", w.id, num)
			}
			w.wg.Done()
		}
	}()
}

// CreateWorkerPool initializes a pool of workers
func CreateWorkerPool(numWorkers int, jobChannel chan Job, wg *sync.WaitGroup) {
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, jobChannel, wg)
		worker.Start()
	}
}

func main() {
	numWorkers := 10
	numJobs := 10 // Total number of jobs (each job processes 10 numbers)

	jobChannel := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Create a worker pool
	CreateWorkerPool(numWorkers, jobChannel, &wg)

	// Create the slice with numbers 1 to 100
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i + 1
	}

	// Send jobs to the worker pool in chunks of 10 numbers
	for i := 0; i < len(data); i += 10 {
		wg.Add(1)
		job := Job{data: data[i : i+10]}
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()

	fmt.Println("All jobs processed")
}
