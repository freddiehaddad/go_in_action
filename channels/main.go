// Demonstration of unbuffered channels.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var numRunners int

func main() {
	numRunners = 4

	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)
	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	b := <-baton
	fmt.Printf("Runner %d has the baton\n", b)

	if b == numRunners {
		fmt.Printf("Race is over!\n")
		wg.Done()
		return
	}

	go Runner(baton)
	baton <- b + 1
}
