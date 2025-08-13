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

	api := r.Group("/api/v1")
	{
		api.GET("/hello", helloHandler.Hello)
		api.GET("/worker/info", workerHandler.GetWorkerInfo)
		api.POST("/worker/task", workerHandler.SubmitTask)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
