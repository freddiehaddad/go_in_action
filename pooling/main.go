package main

import (
	"io"
	"log"
	"math/rand"
	"pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	poolSize   = 4
	numThreads = 20
)

var resourceId int32

// Factory function provided to the resource pool manager for creating
// resources on demand.
func resourceFactory() (io.Closer, error) {
	id := atomic.AddInt32(&resourceId, 1)
	log.Println("resource:", id, "created")
	return &resource{
		id: id,
	}, nil
}

// io.Closer interface implementation for the resource.
func (r *resource) Close() error {
	log.Println("resource:", r.id, "closed")
	return nil
}

// Descriptor for the resource created by the createResource factory.
type resource struct {
	id int32
}

// Worker function that simulates acquiring a resource, using it, and returning
// it back to the pool.
func worker(id int, p *pool.Pool) {
	log.Println("worker:", id, "running")
	r, err := p.Acquire()
	if err != nil {
		log.Println("worker:", "failed to acquire resource", err)
		return
	}
	defer p.Release(r)

	// Simulate job running time by sleeping the thread for a short duration.
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	log.Println("worker:", id, "resource:", r.(*resource).id, "finished")
}

// Program entry point.  Creates a resource pool and spawns several go routines
// to simulate acquiring, using, and returning the resource.
func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().Unix())
	wg.Add(numThreads)

	p, _ := pool.New(resourceFactory, poolSize)
	for i := 1; i <= numThreads; i++ {
		// The id is provided to the go routine to avoid using a closure and
		// sharing the variable across all go routines.
		go func(id int) {
			// We setup random sleep durations to simulate the arrival of
			// new work.
			time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
			worker(id, p)
			wg.Done()
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Close the pool
	p.Close()

	log.Println("main:", "Shutting down")

	// Log number of resources created for statistical data.
	log.Println("main:", "Resources created", resourceId)
}
