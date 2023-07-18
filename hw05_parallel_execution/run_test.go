package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})

	t.Run("concurrency without timers", func(t *testing.T) {
		tasksCount := 6000
		tasks := make([]Task, 0, tasksCount)
		result := make([]int, 0, tasksCount)
		expected := make([]int, 0, tasksCount)
		var mu sync.Mutex
		for i := 0; i < tasksCount; i++ {
			expected = append(expected, i)
			i := i
			tasks = append(tasks, func() error {
				mu.Lock()
				result = append(result, i)
				mu.Unlock()
				return nil
			})
		}

		workersCount := 20
		maxErrorsCount := 5

		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)
		require.NotEqual(t, result, expected, "not all tasks were completed")
	})

	t.Run("task with errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)
		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			if i > 10 {
				err = nil
			}
			tasks = append(tasks, func() error {
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 12

		errorr := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, errorr)
	})

	t.Run("errors := - 1", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)
		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			if i > 10 {
				err = nil
			}
			tasks = append(tasks, func() error {
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := -1

		errorr := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, errorr)
	})

}
