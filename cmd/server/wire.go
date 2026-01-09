//go:build wireinject
// +build wireinject

package main

import (
	"gin-hello-world/internal/handler"
	"gin-hello-world/internal/service"

	"github.com/google/wire"
)

// Handlers 聚合所有 handler
type Handlers struct {
	HelloHandler     *handler.HelloHandler
	WorkerHandler    *handler.WorkerHandler
	SnowflakeHandler *handler.SnowflakeHandler
}

// ServiceSet 定义 service 层的 providers
var ServiceSet = wire.NewSet(
	service.NewHelloService,
	service.NewWorkerService,
)

// HandlerSet 定义 handler 层的 providers
var HandlerSet = wire.NewSet(
	handler.NewHelloHandler,
	handler.NewWorkerHandler,
	handler.NewSnowflakeHandler,
)

// InitializeHandlers 初始化所有 handlers
func InitializeHandlers() (*Handlers, error) {
	wire.Build(
		ServiceSet,
		HandlerSet,
		wire.Struct(new(Handlers), "*"),
	)
	return nil, nil
}
