package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gin-hello-world/pkg/worker"

	"github.com/gin-gonic/gin"
)

// SnowflakeHandler 雪花ID线程池处理器
type SnowflakeHandler struct {
	threadPool worker.ThreadPool
	logger     *worker.Logger
}

// NewSnowflakeHandler 创建新的雪花ID处理器
func NewSnowflakeHandler() *SnowflakeHandler {
	return &SnowflakeHandler{
		logger: worker.NewLogger("SnowflakeHandler"),
	}
}

// StartThreadPool 启动线程池
func (h *SnowflakeHandler) StartThreadPool(c *gin.Context) {
	// 解析参数
	workerCountStr := c.DefaultQuery("workers", "4")
	workerCount, err := strconv.Atoi(workerCountStr)
	if err != nil || workerCount <= 0 {
		h.logger.Error("Invalid worker count: %s", workerCountStr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid worker count",
			"message": "Worker count must be a positive integer",
		})
		return
	}

	machineIDStr := c.DefaultQuery("machine_id", "1")
	machineID, err := strconv.ParseInt(machineIDStr, 10, 64)
	if err != nil || machineID < 0 || machineID > 1023 {
		h.logger.Error("Invalid machine ID: %s", machineIDStr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid machine ID",
			"message": "Machine ID must be between 0 and 1023",
		})
		return
	}

	// 如果已经有运行的线程池，先停止
	if h.threadPool != nil {
		h.logger.Info("Stopping existing thread pool")
		h.threadPool.Stop()
	}

	// 创建新的线程池
	h.threadPool = worker.NewSnowflakeThreadPool(workerCount, machineID)

	// 启动线程池
	ctx := context.Background()
	if err := h.threadPool.Start(ctx); err != nil {
		h.logger.Error("Failed to start thread pool: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to start thread pool",
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("Thread pool started successfully with %d workers, machine ID: %d",
		workerCount, machineID)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Thread pool started successfully",
		"workers":    workerCount,
		"machine_id": machineID,
		"status":     h.threadPool.GetStatus(),
	})
}

// StartSnowflakeLogging 启动雪花ID日志记录
func (h *SnowflakeHandler) StartSnowflakeLogging(c *gin.Context) {
	if h.threadPool == nil {
		h.logger.Error("Thread pool not initialized")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Thread pool not started",
			"message": "Please start the thread pool first",
		})
		return
	}

	// 解析间隔参数
	intervalStr := c.DefaultQuery("interval", "2s")
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		h.logger.Error("Invalid interval format: %s", intervalStr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid interval format",
			"message": "Use format like '1s', '500ms', '2m'",
		})
		return
	}

	// 启动雪花ID日志记录
	ctx := context.Background()
	if err := h.threadPool.StartSnowflakeLogging(ctx, interval); err != nil {
		h.logger.Error("Failed to start snowflake logging: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to start snowflake logging",
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("Snowflake logging started with interval: %v", interval)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Snowflake logging started successfully",
		"interval": interval.String(),
		"status":   h.threadPool.GetStatus(),
	})
}

// GetThreadPoolStatus 获取线程池状态
func (h *SnowflakeHandler) GetThreadPoolStatus(c *gin.Context) {
	if h.threadPool == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Thread pool not initialized",
			"status":  nil,
		})
		return
	}

	status := h.threadPool.GetStatus()
	c.JSON(http.StatusOK, gin.H{
		"message": "Thread pool status retrieved successfully",
		"status":  status,
	})
}

// StopThreadPool 停止线程池
func (h *SnowflakeHandler) StopThreadPool(c *gin.Context) {
	if h.threadPool == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Thread pool not initialized",
			"message": "No thread pool to stop",
		})
		return
	}

	if err := h.threadPool.Stop(); err != nil {
		h.logger.Error("Failed to stop thread pool: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to stop thread pool",
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("Thread pool stopped successfully")
	h.threadPool = nil

	c.JSON(http.StatusOK, gin.H{
		"message": "Thread pool stopped successfully",
	})
}

// SubmitCustomTask 提交自定义任务
func (h *SnowflakeHandler) SubmitCustomTask(c *gin.Context) {
	if h.threadPool == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Thread pool not initialized",
			"message": "Please start the thread pool first",
		})
		return
	}

	var request struct {
		Message string `json:"message"`
		Count   int    `json:"count"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	if request.Count <= 0 {
		request.Count = 1
	}
	if request.Message == "" {
		request.Message = "Custom task"
	}

	// 提交任务
	for i := 0; i < request.Count; i++ {
		taskID := i + 1
		task := func(id int, msg string) func() {
			return func() {
				h.logger.Info("Executing custom task %d: %s", id, msg)
				time.Sleep(100 * time.Millisecond) // 模拟任务执行时间
			}
		}(taskID, request.Message)

		if err := h.threadPool.Submit(task); err != nil {
			h.logger.Error("Failed to submit task %d: %v", taskID, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Tasks submitted successfully",
		"task_count":   request.Count,
		"task_message": request.Message,
		"status":       h.threadPool.GetStatus(),
	})
}
