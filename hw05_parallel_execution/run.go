package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tch := make(chan Task, len(tasks))
	wg := &sync.WaitGroup{}
	wg.Add(n)
	errors := int32(0)

	if n == 0 {
		return nil
	}

	for _, t := range tasks {
		tch <- t
	}
	close(tch)

	for i := 0; i < n; i++ {
		go func(err int32) {
			defer wg.Done()
			for tFunc := range tch {
				if err >= 0 && atomic.LoadInt32(&errors) >= err {
					break
				}
				if result := tFunc(); result != nil {
					atomic.AddInt32(&errors, 1)
				}
			}
		}(int32(m))
	}

	wg.Wait()
	if m >= 0 && errors >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
