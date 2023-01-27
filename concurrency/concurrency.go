package main

import (
	"fmt"
	"os"
	"runner"
	"time"
)

const timeout = 5 * time.Second

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processing task #%d.\n", id)
		time.Sleep(time.Duration(time.Second))
	}
}

func main() {
	timeout := time.Duration(timeout)
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask(), createTask())
	fmt.Printf("Runner initialize %+v\n", r)

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			fmt.Printf("Terminatign due to interrupt\n")
			os.Exit(1)
		case runner.ErrTimeout:
			fmt.Printf("Terminating due to timeout\n")
			os.Exit(2)
		}
	}
	fmt.Printf("Runner completed\n")
}
