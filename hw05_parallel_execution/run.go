package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Counter struct {
	mu sync.RWMutex
	v  int
}

func (counter *Counter) Store() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.v++
}

func (counter *Counter) Load() int {
	counter.mu.RLock()
	defer counter.mu.RUnlock()
	v := counter.v
	return v
}

func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	defer wg.Wait()
	counter := Counter{}
	workers := make(chan string, n)
	defer close(workers)

	for _, task := range tasks {
		if counter.Load() >= m {
			return ErrErrorsLimitExceeded
		}
		wg.Add(1)
		workers <- ""
		go func(task Task, counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			err := task()
			if err != nil {
				counter.Store()
			}
			<-workers
		}(task, &counter, &wg)
	}
	return nil
}
