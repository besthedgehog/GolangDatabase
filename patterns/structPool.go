package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	jobs chan int
	wg sync.WaitGroup
}

func (p *WorkerPool) worker(id int) {
	defer p.wg.Done()
	for job := range p.jobs {
		fmt.Printf("Worker %d: job %d\n", id, job)
	}
}

func NewWorkerPool(numberOfWorkers int) *WorkerPool {
	pool := &WorkerPool{
		jobs: make(chan int),
		// wg не создаём?
	}
	pool.wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		go pool.worker(i)
	}
	return pool
}

func (p *WorkerPool) Submit(job int) {
	p.jobs <- job
}

func (p *WorkerPool) Close() {
	close(p.jobs)
	p.wg.Wait()
}

func main() {
	pool := NewWorkerPool(3)
	for j := range 10 {
		pool.Submit(j)
	}
	pool.Close()
}
