package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded  = errors.New("errors limit exceeded")
	ErrTooLowGorutineNumber = errors.New("at least two gorutine must be specified for concurrently execution")
)

type (
	Task         func() error
	errorCounter struct {
		mu    sync.RWMutex
		count int
	}
)

func (c *errorCounter) add() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *errorCounter) get() int {
	c.mu.RLock()
	count := c.count
	c.mu.RUnlock()
	return count
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 2 {
		return ErrTooLowGorutineNumber
	}
	chErr := make(chan error)
	chTask := make(chan Task)
	chStop := make(chan struct{})
	wg := sync.WaitGroup{}
	errCount := errorCounter{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go executor(chErr, chTask, &wg)
	}

	go errIncreaser(chErr, chStop, &errCount)

	for _, task := range tasks {
		if errCount.get() >= m && m > 0 {
			// stop iteration on tasks cause of errors
			break
		}
		chTask <- task
	}

	close(chTask)
	wg.Wait()
	close(chStop)

	// ignore all errors, when m <= 0
	if errCount.get() >= m && m > 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func errIncreaser(chErr <-chan error, chStop <-chan struct{}, errCount *errorCounter) {
	for {
		select {
		case <-chStop:
			// stop errIncreaser
			return
		case <-chErr:
			errCount.add()
		default:
			continue
		}
	}
}

func executor(chErr chan<- error, chTask <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-chTask
		if !ok {
			// stop executor when closed chTask
			return
		}

		err := task()
		if err != nil {
			// sending error to channel chErr
			chErr <- err
		}
	}
}
