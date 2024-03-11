package worker

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type Context struct {
	Ctx      context.Context
	Cancel   context.CancelFunc
	ErrGroup *errgroup.Group
}

func NewWorkerContext() *Context {
	ctx, cancel := context.WithCancel(context.Background())
	errGroup, ctx := errgroup.WithContext(ctx)
	return &Context{
		Ctx:      ctx,
		Cancel:   cancel,
		ErrGroup: errGroup,
	}
}

func (ctx *Context) WaitAndClose() error {
	err := ctx.ErrGroup.Wait()
	ctx.Cancel()

	return err
}
