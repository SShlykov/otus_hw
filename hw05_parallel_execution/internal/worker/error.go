package worker

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

func ErrorHandler(context *Context, errChan <-chan error, errorLimit int) error {
	var errorsCount int
	for {
		select {
		case <-context.Ctx.Done():
			return nil
		case <-errChan:
			errorsCount++
			if errorsCount >= errorLimit {
				context.Cancel()
				return ErrErrorsLimitExceeded
			}
		}
	}
}
