package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrorsCount struct {
	count int32
}

func (e *ErrorsCount) Get() int32 {
	return atomic.LoadInt32(&e.count)
}

func (e *ErrorsCount) Increment() {
	atomic.AddInt32(&e.count, 1)
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	var wg sync.WaitGroup
	ch := make(chan Task)
	e := &ErrorsCount{}

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					e.Increment()
				}
			}
		}()
	}

	for _, task := range tasks {
		ch <- task
		if int(e.Get()) >= m {
			break
		}
	}

	close(ch)
	wg.Wait()

	if e.Get() > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
