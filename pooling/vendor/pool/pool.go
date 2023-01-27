// A package for syncronized management of shared resources
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool closed error.
var ErrPoolClosed = errors.New("Pool closed")

// Manger of a pool of resources.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// Creates a pool
func New(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Invalid size")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   factory,
	}, nil
}

// Return a resource from the pool or a new one.
func (p *Pool) Acquire() (io.Closer, error) {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release a resource adding it back to the pool or closing it if the pool is full.
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		log.Println("Release:", "Pool Closed", "Closing")
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
		return

	default:
		r.Close()
		log.Println("Release:", "Closing")
		return
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		log.Println("Close:", "Closed")
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		log.Println("Close:", "Closing")
		r.Close()
	}
}
