package handler

import (
	"gin-hello-world/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	helloService *service.HelloService
}

func NewHelloHandler(helloService *service.HelloService) *HelloHandler {
	return &HelloHandler{
		helloService: helloService,
	}
}

func (h *HelloHandler) Hello(c *gin.Context) {
	message := h.helloService.GetHelloMessage()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    message,
	})
}
