package main

import (
	"gin-hello-world/internal/handler"
	"gin-hello-world/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	helloService := service.NewHelloService()
	helloHandler := handler.NewHelloHandler(helloService)

	workerService := service.NewWorkerService()
	workerHandler := handler.NewWorkerHandler(workerService)

	snowflakeHandler := handler.NewSnowflakeHandler()

	api := r.Group("/api/v1")
	{
		api.GET("/hello", helloHandler.Hello)
		api.GET("/worker/info", workerHandler.GetWorkerInfo)
		api.POST("/worker/task", workerHandler.SubmitTask)

		// 雪花ID线程池相关接口
		api.POST("/snowflake/start", snowflakeHandler.StartThreadPool)
		api.POST("/snowflake/logging/start", snowflakeHandler.StartSnowflakeLogging)
		api.GET("/snowflake/status", snowflakeHandler.GetThreadPoolStatus)
		api.POST("/snowflake/stop", snowflakeHandler.StopThreadPool)
		api.POST("/snowflake/task", snowflakeHandler.SubmitCustomTask)
	}

	log.Println("Server starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
