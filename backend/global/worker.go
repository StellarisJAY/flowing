package global

import (
	"flowing/internal/config"
	"runtime"

	"github.com/panjf2000/ants/v2"
)

var (
	workerPool *ants.Pool
)

func InitWorkerPool(config *config.Config) {
	pool, err := ants.NewPool(min(max(0, config.GoPool.Size), runtime.NumCPU()))
	if err != nil {
		panic(err)
	}
	workerPool = pool
}

func WorkerPool() *ants.Pool {
	return workerPool
}
