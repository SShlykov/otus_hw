package hw05parallelexecution

import (
	"errors"
	"github.com/SShlykov/otus_hw/hw05_parallel_execution/internal/worker"
	"sync"
)

// Run starts tasks in workerCount goroutines and stops its work when receiving errorLimit errors from tasks.
func Run(tasks []worker.Task, workerCount, errorLimit int) error {
	if workerCount < 0 {
		return errors.New("workerCount must be greater than 0")
	}
	context := worker.NewWorkerContext()
	wg := &sync.WaitGroup{}

	errChan := make(chan error, workerCount)
	doneChan := make(chan struct{}, workerCount)
	taskChan := make(chan worker.Task)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.Start(context.Ctx, taskChan, doneChan, errChan)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, task := range tasks {
			select {
			case taskChan <- task:
			case <-context.Ctx.Done():
				return
			}
		}
	}()

	context.ErrGroup.Go(func() error { return worker.ErrorHandler(context, errChan, errorLimit) })

	go worker.Finisher(context, len(tasks), doneChan)

	err := context.WaitAndClose()
	wg.Wait()

	return err
}
