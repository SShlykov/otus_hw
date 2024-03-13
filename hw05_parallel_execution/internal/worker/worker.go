package worker

import (
	"context"
)

type Task func() error

func Start(ctx context.Context, taskChan <-chan Task, doneChan chan<- struct{}, errChan chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-taskChan:
			if task == nil || ctx.Err() != nil {
				break
			}
			handleTaskResult(task, errChan, doneChan)
		}
	}
}

func handleTaskResult(task Task, errChan chan<- error, doneChan chan<- struct{}) {
	err := task()
	if err != nil {
		errChan <- err
	}
	doneChan <- struct{}{}
}
