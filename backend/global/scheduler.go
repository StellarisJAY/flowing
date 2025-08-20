package global

import (
	"flowing/internal/config"
	"flowing/internal/scheduler"
	"flowing/internal/scheduler/pool"
)

var worker scheduler.Worker

func InitScheduler(config *config.Config) {
	var err error
	switch config.Worker {
	case "gopool":
		worker, err = pool.NewGoPoolWorker(config.GoPool.Size, config.GoPool.Timeout)
	default:
		panic("invalid worker " + config.Worker)
	}
	if err != nil {
		panic(err)
	}
}

func Worker() scheduler.Worker {
	return worker
}
