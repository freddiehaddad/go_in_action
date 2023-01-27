// Package runner manages the running and lifetime of a process.
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be stopped using an operating system inturrupt.
type Runner struct {
	// interrupt channel reports a signal from the operating system.
	interrupt chan os.Signal

	// complete channel reports that processing is done.
	complete chan error

	// timeout reports that time has run out.
	timeout <-chan time.Time

	// tasks holds a set of functions that are executed synchronously in index order.
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout.
var ErrTimeout = errors.New("timeout received")

// ErrInterrupt is returned when an operating system inturrupt is received.
var ErrInterrupt = errors.New("interrupt received")

// Returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Adds the variable list of functions to the task holder.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Starts the Runner.
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// Tasks completed
	case err := <-r.complete:
		return err

	// interrupt received
	case <-r.timeout:
		return ErrTimeout
	}
}

// Internal task runner.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// Checks if an interrupt was received.
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
