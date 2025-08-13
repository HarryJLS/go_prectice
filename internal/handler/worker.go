package handler

import (
	"gin-hello-world/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkerHandler struct {
	workerService *service.WorkerService
}

func NewWorkerHandler(workerService *service.WorkerService) *WorkerHandler {
	return &WorkerHandler{
		workerService: workerService,
	}
}

func (h *WorkerHandler) GetWorkerInfo(c *gin.Context) {
	info := h.workerService.GetWorkerInfo()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info,
	})
}

func (h *WorkerHandler) SubmitTask(c *gin.Context) {
	h.workerService.SubmitTask()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Task submitted successfully",
		"data":    nil,
	})
}