package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tch := make(chan Task, len(tasks))
	errch := make(chan error, len(tasks))
	wg := &sync.WaitGroup{}
	wg.Add(n)

	var errors error = nil

	for _, t := range tasks {
		fmt.Println("write task")
		tch <- t
	}
	fmt.Println("write completed!")

	for i := 0; i < n; i++ {
		go func(tasks chan Task, errch chan error) {
			defer wg.Done()
			for {
				tFunc, opened := <-tch
				if opened {
					if err := tFunc(); err != nil {
						errch <- ErrErrorsLimitExceeded
						break
					}
					continue
				}
				break
			}
		}(tch, errch)
	}

	fmt.Println("wait error")
	if errors = <-errch; errors != nil {
		fmt.Println("detected")
		close(tch)
	}
	wg.Wait()
	return errors
}
