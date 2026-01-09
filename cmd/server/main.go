package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用 wire 初始化所有依赖
	handlers, err := InitializeHandlers()
	if err != nil {
		log.Fatal("Failed to initialize handlers:", err)
	}

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/hello", handlers.HelloHandler.Hello)
		api.GET("/worker/info", handlers.WorkerHandler.GetWorkerInfo)
		api.POST("/worker/task", handlers.WorkerHandler.SubmitTask)

		// 雪花ID线程池相关接口
		api.POST("/snowflake/start", handlers.SnowflakeHandler.StartThreadPool)
		api.POST("/snowflake/logging/start", handlers.SnowflakeHandler.StartSnowflakeLogging)
		api.GET("/snowflake/status", handlers.SnowflakeHandler.GetThreadPoolStatus)
		api.POST("/snowflake/stop", handlers.SnowflakeHandler.StopThreadPool)
		api.POST("/snowflake/task", handlers.SnowflakeHandler.SubmitCustomTask)
	}

	log.Println("Server starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
