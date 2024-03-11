package hw05parallelexecution

import (
	"github.com/SShlykov/otus_hw/hw05_parallel_execution/internal/worker"
	"sync"
)

// Run starts tasks in workerCount goroutines and stops its work when receiving errorLimit errors from tasks.
func Run(tasks []worker.Task, workerCount, errorLimit int) error {
	context := worker.NewWorkerContext()
	wg := &sync.WaitGroup{}

	errChan := make(chan error, errorLimit)
	doneChan := make(chan struct{}, len(tasks))
	taskChan := initTaskChan(tasks)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker.Start(context.Ctx, wg, taskChan, doneChan, errChan)
	}

	context.ErrGroup.Go(func() error { return worker.ErrorHandler(context, errChan, errorLimit) })

	go worker.Finisher(context, len(tasks), doneChan)

	err := context.WaitAndClose()
	wg.Wait()

	return err
}

func initTaskChan(tasks []worker.Task) <-chan worker.Task {
	taskChan := make(chan worker.Task, len(tasks))
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	return taskChan
}
