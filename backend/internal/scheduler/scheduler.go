package scheduler

import "context"

type Task interface {
	Run(ctx context.Context) error
	ID() string
}

type Worker interface {
	Submit(task Task) error
	Shutdown() error
	Cancel(id string)
}
