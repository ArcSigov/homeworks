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
	tch := make(chan Task)
	wg := &sync.WaitGroup{}
	wg.Add(n + 1)

	var errors int32 = 0

	go func() {
		defer wg.Done()
		for _, t := range tasks {
			if atomic.LoadInt32(&errors) >= int32(m) {
				break
			}
			tch <- t
		}
		close(tch)
	}()

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				if tFunc, ok := <-tch; ok {
					if err := tFunc(); err != nil {
						atomic.AddInt32(&errors, 1)
					}
				} else {
					break
				}
			}
		}()
	}

	wg.Wait()
	if errors >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
