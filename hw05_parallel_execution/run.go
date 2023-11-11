package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded  = errors.New("errors limit exceeded")
	ErrTooLowGorutineNumber = errors.New("at least two gorutine must be specified for concurrently execution")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 2 {
		return ErrTooLowGorutineNumber
	}
	var errCount int32
	chTask := make(chan Task)
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go executor(&errCount, chTask, &wg)
	}

	for _, task := range tasks {
		if int(atomic.LoadInt32(&errCount)) >= m && m > 0 {
			// stop iteration on tasks cause of errors
			break
		}
		chTask <- task
	}

	close(chTask)
	wg.Wait()

	// ignore all errors, when m <= 0
	if int(atomic.LoadInt32(&errCount)) >= m && m > 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func executor(errCount *int32, chTask <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range chTask {
		err := task()
		if err != nil {
			// increase errors
			atomic.AddInt32(errCount, 1)
		}
	}
}
