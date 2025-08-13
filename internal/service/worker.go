package service

import (
	"gin-hello-world/pkg/worker"
	"math/rand"
	"time"
)

type WorkerService struct {
	pool *worker.WorkerPool
}

func NewWorkerService() *WorkerService {
	pool := worker.NewWorkerPool(5)
	pool.Start()
	
	return &WorkerService{
		pool: pool,
	}
}

func (s *WorkerService) GetWorkerInfo() *worker.PoolInfo {
	return s.pool.GetPoolInfo()
}

func (s *WorkerService) SubmitTask() {
	s.pool.Submit(func() {
		duration := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(duration)
	})
}

func (s *WorkerService) Stop() {
	s.pool.Stop()
}