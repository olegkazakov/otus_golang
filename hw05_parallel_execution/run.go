package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(tasks <-chan Task, errors chan<- error) {
	for task := range tasks {
		if err := task(); err != nil {
			select {
			case errors <- err:
			default:
				return
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	if n <= 0 || len(tasks) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(n)

	taskChannel := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel)

	errorChannel := make(chan error, m)
	defer close(errorChannel)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			worker(taskChannel, errorChannel)
		}()
	}

	wg.Wait()

	if len(errorChannel) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
