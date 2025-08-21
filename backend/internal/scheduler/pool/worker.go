package pool

import (
	"context"
	"flowing/internal/scheduler"
	"log/slog"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

type GoPoolWorker struct {
	pool    *ants.Pool
	timeout time.Duration
	tasks   sync.Map
}

func NewGoPoolWorker(size int, timeout time.Duration) (*GoPoolWorker, error) {
	pool, err := ants.NewPool(size)
	if err != nil {
		return nil, err
	}
	return &GoPoolWorker{pool: pool, timeout: timeout, tasks: sync.Map{}}, nil
}

func (g *GoPoolWorker) Submit(task scheduler.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	g.tasks.Store(task.ID(), cancel)
	return g.pool.Submit(func() {
		defer func() {
			cancel()
			if err := recover(); err != nil {
				slog.Error("go pool worker panic", "err", err)
			}
		}()
		if err := task.Run(ctx); err != nil {
			slog.Error("go pool worker run task error", "err", err)
		}
		slog.Info("go pool worker run task success", "task_id", task.ID())
	})
}

func (g *GoPoolWorker) Shutdown() error {
	g.pool.Release()
	return nil
}

func (g *GoPoolWorker) Cancel(id string) {
	if cancel, ok := g.tasks.Load(id); ok {
		cancel.(context.CancelFunc)()
		g.tasks.Delete(id)
	}
}
