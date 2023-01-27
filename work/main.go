// Driver to demonstrate work pool.
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
	"work"
)

const (
	maxGoroutines = 11
	numTasks      = 100
)

type task struct {
	id int
}

func (t *task) Task() {
	log.Println("Task: ", t.id, "started")
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	log.Println("Task: ", t.id, "finished")
}

func main() {
	log.SetFlags(log.LUTC | log.Ldate | log.Lmicroseconds | log.Lshortfile)

	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup

	p := work.New(maxGoroutines)

	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		t := task{i}
		go func() {
			p.Run(&t)
			wg.Done()
		}()
	}

	log.Println("Waiting for task creation to finish...")
	wg.Wait()

	log.Println("Shutting down...")
	p.Shutdown()
	log.Println("Done")
}
