// A demonstration of buffered channels.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers     = 4
	bufferCapacity = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	channel := make(chan string, bufferCapacity)
	wg.Add(numWorkers)

	// start workers
	for nw := 1; nw <= numWorkers; nw++ {
		fmt.Printf("Creating worker %d ...\n", nw)
		go worker(channel, nw)
	}

	// create work
	for job := 1; job <= bufferCapacity; job++ {
		channel <- fmt.Sprintf("Task: %d", job)
	}
	close(channel)

	wg.Wait()
	fmt.Printf("All workers finished, exiting.\n")
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	var maxSleep int64
	maxSleep = 100

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d shutting down. No more work.\n", worker)
			return
		}

		fmt.Printf("Worker %d processing task %s ...\n", worker, task)
		sleep := rand.Int63n(maxSleep)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker %d finished processing task %s.\n", worker, task)
	}
}
